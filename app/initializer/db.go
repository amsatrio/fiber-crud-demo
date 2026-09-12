package initializer

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB_GERMAN *gorm.DB
var DB_HOSPITAL *gorm.DB

func getGormLogLevel() logger.LogLevel {
	logMode := os.Getenv("LOG_MODE")
	switch logMode {
	case "ERROR":
		return logger.Error
	case "INFO":
		return logger.Info
	default:
		return logger.Info
	}
}

func getGormWriter(filePath string) (io.Writer, *os.File) {
	if os.Getenv("LOG_TYPE") == "stdout" {
		return os.Stdout, nil
	}

	if err := os.MkdirAll(filepath.Dir(filePath), 0755); err != nil {
		panic(err)
	}

	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}

	return io.MultiWriter(os.Stdout, file), file
}

func newGormLogger(filePath string) (*logger.Interface, io.Closer) {
	writer, file := getGormWriter(filePath)

	l := logger.New(
		log.New(writer, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  getGormLogLevel(),
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      true,
			Colorful:                  false,
		},
	)

	var closer io.Closer
	if file != nil {
		closer = file
	}

	return &l, closer
}

// ensureDSNTimeout injects a default dial timeout into a MySQL DSN if the DSN
// does not already specify one, so that connecting to an unreachable host
// fails fast instead of hanging the serverless invocation.
func ensureDSNTimeout(dsn string) string {
	withTimeout := func(prefix string) string {
		if strings.Contains(dsn, prefix+"=") {
			return dsn
		}
		sep := "?"
		if strings.Contains(dsn, "?") {
			sep = "&"
		}
		return dsn + sep + prefix + "=5s"
	}
	dsn = withTimeout("timeout")
	dsn = withTimeout("readTimeout")
	dsn = withTimeout("writeTimeout")
	return dsn
}

func openDB(dsnEnvKey, logFile string) (*gorm.DB, error) {
	gormLogger, closer := newGormLogger(logFile)
	if closer != nil {
		defer closer.Close()
	}

	dsn := ensureDSNTimeout(os.Getenv(dsnEnvKey))
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // data source name
		DefaultStringSize:         256,   // default size for string fields
		DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,  // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
	}), &gorm.Config{
		Logger: *gormLogger,
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Minute)

	return db, nil
}

func InitializeDatabaseGerman() {
	db, err := openDB("DB_GERMAN_URL", "logs/gorm-db-german.log")
	if err != nil {
		// Don't kill the process: on serverless, a failed DB connection for one
		// module must not crash the whole function (which surfaces as
		// FUNCTION_INVOCATION_FAILED). Log and continue.
		log.Printf("error: failed to connect to DB: %v\n", err)
		return
	}
	DB_GERMAN = db
}

func InitializeDatabaseHospital() {
	db, err := openDB("DB_HOSPITAL_URL", "logs/gorm-db-hospital.log")
	if err != nil {
		log.Printf("error: failed to connect to DB_HOSPITAL: %v\n", err)
		return
	}
	DB_HOSPITAL = db
}
