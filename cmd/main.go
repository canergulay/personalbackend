package main

import (
	"github.canergulay/blogbackend/internal/database"
	"github.canergulay/blogbackend/internal/server"
	"github.canergulay/blogbackend/internal/server/routes/blog"
	"github.canergulay/blogbackend/internal/server/routes/home"
	"github.canergulay/blogbackend/internal/server/services"
	"github.canergulay/blogbackend/internal/server/websocket"
	"github.canergulay/blogbackend/internal/server/websocket/handlers"
)

func main() {

	pgManager := database.InitPG()
	blogManager := blog.NewBlogManager(pgManager.DB)
	createPostService := services.NewCreatePostService(blogManager)
	savePostService := services.NewSavePostService(blogManager)
	getPostsService := services.NewGetPostService(blogManager)
	websocketHandler := handlers.NewWebSocketHandler(&createPostService, &savePostService, &getPostsService)
	sv := server.InitialiseAllRoutes(blog.NewBlogManager(pgManager.DB), websocket.NewSocketManager(&websocketHandler), home.NewHomeManager())
	sv.StartServer(":8080")

}
