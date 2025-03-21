package main

import (
	"go-jwt-api/config"
	"go-jwt-api/routes"
)

func main() {
	config.ConnectDatabase()
	r := routes.SetupRouter()
	r.Run(":8080")
}
