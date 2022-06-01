package main

import (
	"github.canergulay/blogbackend/internal/database"
	"github.canergulay/blogbackend/internal/server"
	"github.canergulay/blogbackend/internal/server/routes/blog"
	"github.canergulay/blogbackend/internal/server/routes/home"
	"github.canergulay/blogbackend/internal/server/services"
	"github.canergulay/blogbackend/internal/server/websocket"
)

func main() {

	pgManager := database.InitPG()
	blogManager := blog.NewBlogManager(pgManager.DB)
	createPostHandler := services.NewCreatePostHandler(&blogManager)
	savePostHandler := services.NewSavePostHandler(&blogManager)
	sv := server.InitialiseAllRoutes(websocket.NewSocketManager(&createPostHandler, &savePostHandler), home.NewHomeManager())
	sv.StartServer(":8080")

}
