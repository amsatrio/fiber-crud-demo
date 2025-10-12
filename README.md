# fiber-crud-demo

## create new project
go mod init fiber-crud-demo

## dependencies
go get github.com/gofiber/fiber/v2
go get github.com/gofiber/fiber/v2/middleware/cache
go get gorm.io/gorm
go get gorm.io/driver/mysql
go get github.com/joho/godotenv
go get github.com/go-playground/validator/v10
go get github.com/goccy/go-json
go get github.com/golang-jwt/jwt/v5
go install github.com/swaggo/swag/cmd/swag@latest
go get github.com/gofiber/template/html/v2