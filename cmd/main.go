package main

import (
	"github.canergulay/blogbackend/internal/database"
	"github.canergulay/blogbackend/internal/server"
	"github.canergulay/blogbackend/internal/server/routes/blog"
	"github.canergulay/blogbackend/internal/server/routes/home"
	"github.canergulay/blogbackend/internal/server/routes/upload"
	"github.canergulay/blogbackend/internal/server/services"
	"github.canergulay/blogbackend/internal/server/websocket"
	"github.canergulay/blogbackend/internal/server/websocket/handlers"
)

func main() {
	pgManager := database.InitPG()
	createPostService := services.NewCreatePostService(pgManager.DB)
	savePostService := services.NewSavePostService(pgManager.DB)
	getPostsService := services.NewGetPostService(pgManager.DB)
	blogManager := blog.NewBlogManager(&savePostService, &getPostsService)
	websocketHandler := handlers.NewWebSocketHandler(&createPostService, &savePostService, &getPostsService)
	sv := server.InitialiseAllRoutes(blogManager, websocket.NewSocketManager(&websocketHandler), home.NewHomeManager(), upload.NewUploadManager())
	sv.StartServer(":8080")

}
