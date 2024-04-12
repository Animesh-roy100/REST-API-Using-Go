package main

import (
	_ "main.go/docs"
	"main.go/initializers"
	"main.go/router"
)

func main() {
	initializers.LoadEnv()
	router.Run()
}
