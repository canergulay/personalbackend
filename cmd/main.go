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
	// DATABASE INITIALIZATION \\
	pgManager := database.InitPG()

	// SERVICES INITIALIZATION \\
	createPostService := services.NewCreatePostService(pgManager.DB)
	savePostService := services.NewSavePostService(pgManager.DB)
	getPostsService := services.NewGetPostService(pgManager.DB)
	getCommentsService := services.NewGetCommentsService(pgManager.DB)

	// AS THE NAME SUGGESTS.. \\
	websocketHandler := handlers.NewWebSocketHandler(&createPostService, &savePostService, &getPostsService)

	// A MANAGER IS LIKE MODULE OF A FEW COMPONENTS. THEY USE SERVICES TO OBTAIN CERTAIN FUNCTIONALITIES. \\
	blogManager := blog.NewBlogManager(&savePostService, &getPostsService, &getCommentsService)
	homeManager := home.NewHomeManager()
	uploadManager := upload.NewUploadManager()
	webSocketManager := websocket.NewSocketManager(&websocketHandler)

	// KICK THE ASS OF THE SERVER \\
	sv := server.InitialiseAllRoutes(blogManager, webSocketManager, homeManager, uploadManager)

	sv.StartServer(":8080")

}
