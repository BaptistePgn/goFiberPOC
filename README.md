# FIBER GO POC
# This is a simple POC to test the fiber framework in go

## How to run
copy the .env.example to .env and fill the database connection string
```bash
docker-compose up -d

go mod init fiber-apis

go get -u github.com/gofiber/fiber/v3
go get github.com/lib/pq
go get github.com/joho/godotenv
go get -u gorm.io/gorm

go run main.go
```
