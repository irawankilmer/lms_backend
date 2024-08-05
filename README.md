# lms_backend
Golang rest API

## Library
1. go get -u github.com/gin-gonic/gin
2. go get -u gorm.io/gorm
3. go get -u gorm.io/driver/postgres
4. go get -u gorm.io/driver/mysql
5. go get github.com/joho/godotenv
6. go get -u github.com/swaggo/swag/cmd/swag
7. go install github.com/swaggo/swag/cmd/swag@latest
8. go get -u github.com/swaggo/files
9. go get -u github.com/swaggo/gin-swagger
10. go get github.com/gorilla/websocket
11. go get github.com/gin-contrib/cors
12. go get -u github.com/swaggo/swag/cmd/swag

## Swag 
```bash
swag init --dir ./cmd/myapp,./internal/app/handler --output ./docs --parseDependency
```
