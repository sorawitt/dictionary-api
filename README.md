# Go Dictionary REST API Example
A RESTful API example for simple todo application with Go

It is a just simple tutorial or example for making simple RESTful API with Go using `Gin`, `MongoDB` and `Hexagonal architecture`

## Installation & Run
```bash
# Download this project
go get github.com/sorawitt/dictionary-api
```

Before running API server, you should set the database config with yours or set the your database config with my values on [config.go](https://github.com/mingrammer/go-todo-rest-api-example/blob/master/config/config.go)
```go
```

```bash
# Build and Run
cd dictionary-api
go build
./dictionary-api

# API Endpoint : http://127.0.0.1:3000
```

## Structure
```
├── app
│   ├── app.go
│   ├── repository
│   │   ├── auth.go
│   ├── handler          // Our API core handlers
│   │   ├── common.go    // Common response functions
│   │   ├── auth.go      // APIs for Authentication
│   │   └── profile.go   // APIs for Profile model
│   └── service
│       └── auth.go     // Service for Authentication
├── config
│   └── config.go        // Configuration
└── main.go
```

## API

#### /auth/signup
* `POST` : Create a new user

## Todo

- [ ] Support basic REST APIs.
- [ ] Support Authentication with user for securing the APIs.
- [ ] Make convenient wrappers for creating API handlers.
- [ ] Write the tests for all APIs.
- [ ] Organize the code with packages
- [ ] Make docs with GoDoc
- [ ] Building a deployment process 
