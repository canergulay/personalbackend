package main

import (
	"github.canergulay/blogbackend/internal/database"
	"github.canergulay/blogbackend/internal/server"
)

func main() {

	database.InitPG()
	sv := server.InitialiseAllRoutes()
	sv.StartServer(":8080")

}
