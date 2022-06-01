package main

import (
	"github.canergulay/blogbackend/internal/database"
	"github.canergulay/blogbackend/internal/server"
	"github.canergulay/blogbackend/internal/server/endpoints/routes/blog"
	"github.canergulay/blogbackend/internal/server/endpoints/routes/home"
	"github.canergulay/blogbackend/internal/server/endpoints/routes/websocket"
	"github.canergulay/blogbackend/internal/server/endpoints/routes/websocket/handlers"
)

func main() {

	pgManager := database.InitPG()
	blogManager := blog.NewBlogManager(pgManager.DB)
	createPostHandler := handlers.NewCreatePostHandler(&blogManager)
	savePostHandler := handlers.NewSavePostHandler(&blogManager)
	sv := server.InitialiseAllRoutes(websocket.NewSocketManager(&createPostHandler, &savePostHandler), home.NewHomeManager())
	sv.StartServer(":8080")

}
