# go-orm-jwt

generate Swagger API 

```bash
export PATH=$(go env GOPATH)/bin:$PATH

swag init -g cmd/server/main.go --output docs/api
```

normal run api

```bash
go run cmd/server/main.go
```