package main

import (
	"github.com/ueverson/WebApiGo/database"
	"github.com/ueverson/WebApiGo/server"
)

func main() {
	database.StartDB()

	server := server.NewServer()

	server.Run()
}
