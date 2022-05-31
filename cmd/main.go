package main

import (
	"github.canergulay/blogbackend/internal/server"
)

func main() {

	sv := server.InitialiseAllRoutes()
	sv.StartServer(":8080")
}
