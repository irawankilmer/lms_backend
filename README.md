# lms_backend
Golang rest API
### Migration
```bash
go run cmd/migrate/main.go
```

### Seeder
```bash
go run cmd/seeder/main.go
```
### Swag init
```bash
swag init -g cmd/server/main.go -o ./docs --parseDependency
```

### go run
```bash
go run cmd/server/main.go
```