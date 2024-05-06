# Api Template (small service)

The project is template for api `Http Web framework`. Develop using `Go` with `Gin` as framwork for API.

#### **Project Structure**

```
├───main.go
├───routes
    └───index.go -> all route on app
├───handlers
├───services
├───middlewares
├───db
    └───config.go -> connect & migrate db
```

## Installation

### Install dependency and start

- install [Go](https://go.dev/doc/install)
- install [Air](https://github.com/cosmtrek/air) library for achieving hot-reloading and auto-compilation in applications
  - `go install github.com/cosmtrek/air@latest`
  - `ari init`
- run `go get` to install dependencies
- run `air` to start app

## Tech Stack

**Server:** Go, Gin

**Orm:** Gorm

**Driver:** Mysql

**Env:** Godotenv

**Hot-reloading:** Air
