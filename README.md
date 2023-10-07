## Initial Project

# Tech STack

- [Fiber](https://docs.gofiber.io/)
- [Gorm](https://gorm.io/docs/)
- [Viper](https://github.com/spf13/viper)
- [JWT](https://pkg.go.dev/github.com/dgrijalva/jwt-go@v3.2.0+incompatible)
- [Postgresql](https://www.postgresql.org/docs/)
 
## Environment Variables

To run this project, you will need to add the following environment variables to your public.env file

`DB_HOST:`

`DB_USER:`

`DB_USER:`

`DB_PASSWORD:`

`DB_DBNAME:`

`DB_PORT:`

`SSL_MODE:`

`SECRET_JWT:`

`APP_PORT:`

# Installation

## Run Local

Clone the project

```bash
git clone https://github.com/Arbyusman/Rest-API-Learn.git
```

Go to the project directory

```bash
cd Rest-API-Learn
```

Install Package

```bash
go mod download
```

Start the server

```bash
go build main.go
```

or

```bash
go run  main.go
```
