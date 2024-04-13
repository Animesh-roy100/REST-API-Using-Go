package main

import (
	_ "main.go/docs"
	"main.go/initializers"
	"main.go/router"
	"main.go/storage"
)

// @title REST API Using Go
// @version 1.0
// @description A RESTful API in Go to understand basics

// @host localhost:3001
func main() {
	initializers.LoadEnv()
	storage.LoadDatabase()
	router.Run()
}
