package main

import (
	"github.canergulay/blogbackend/internal/server"
	"github.canergulay/blogbackend/internal/server/endpoints/routes/home"
)

func main() {
	home := home.GetHomeEndpoint()
	sv := server.GetNewServerManager(home)
	sv.StartServer(":8080")
}
