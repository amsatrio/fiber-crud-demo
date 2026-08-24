package initializer

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB
var DB_FILE_MANAGEMENT *gorm.DB

func InitializeDatabase() {
	var err error

	filePath := "logs/gorm-db-hospital.log"

	// 1. Ensure directory exists
	if err := os.MkdirAll(filepath.Dir(filePath), 0755); err != nil {
		panic(err)
	}

	// 2. Create the file
	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}

	multiOutput := io.MultiWriter(os.Stdout, file)

	multiLogger := logger.New(
		log.New(multiOutput, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,        // Don't include params in the SQL log
			Colorful:                  false,       // Disable color
		},
	)

	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // data source name
		DefaultStringSize:         256,   // default size for string fields
		DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,  // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
	}), &gorm.Config{
		Logger: multiLogger,
	})
	if err != nil {
		log.Fatal("Failed to connnect to database")
	}

	sqlDB, err := DB.DB()

	if err != nil {
		log.Fatal("Failed to get sqlDB")
	}

	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Minute)

	defer file.Close()
}

func InitializeDatabaseFileManagement() {
	var err error

	filePath := "logs/gorm-db-file-management.log"

	// 1. Ensure directory exists
	if err := os.MkdirAll(filepath.Dir(filePath), 0755); err != nil {
		panic(err)
	}

	// 2. Create the file
	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}

	multiOutput := io.MultiWriter(os.Stdout, file)

	multiLogger := logger.New(
		log.New(multiOutput, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,        // Don't include params in the SQL log
			Colorful:                  false,       // Disable color
		},
	)

	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := os.Getenv("DB_FILE_MANAGEMENT_URL")
	DB_FILE_MANAGEMENT, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // data source name
		DefaultStringSize:         256,   // default size for string fields
		DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,  // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
	}), &gorm.Config{
		Logger: multiLogger,
	})
	if err != nil {
		log.Fatal("Failed to connnect to database")
	}

	sqlDB, err := DB_FILE_MANAGEMENT.DB()

	if err != nil {
		log.Fatal("Failed to get sqlDB")
	}

	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Minute)

	defer file.Close()
}
