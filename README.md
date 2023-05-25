# go-api
golang rest API

## installation
- go get github.com/gofiber/fiber/v2
- go get github.com/joho/godotenv
- go get -u gorm.io/gorm
- go get -u gorm.io/driver/postgres
- go get -u gorm.io/driver/sqlite


## RUN
- make sure 5432 port is free on your ubunt: 'sudo lsof -i -P -n | grep LISTEN'
- start a postgres instance: docker run --name postgres_db  -p 5432:5432 -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=admin123123! -e POSTGRES_DB=postgres_db -d postgres
- run APP: go run main.go