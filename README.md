# go-tut-RestApi
Tutorial: Developing a RESTful API with Go and Gin
In tutorial we make rest api web app with Gin framework and in memory DB with clices
# Prerequisites
- Go 1.16 or later.
- The curl tool.

# Endpoints
### /albums
- GET – Get a list of all albums, returned as JSON.
- POST – Add a new album from request data sent as JSON.
### /albums/:id
- GET – Get an album by its ID, returning the album data as JSON.

# Dependencies
- [Gin framework](https://github.com/gin-gonic/gin)

# Run
## get deps
```sh
go mod tidy
```
## run
```sh
go run main.go
```
