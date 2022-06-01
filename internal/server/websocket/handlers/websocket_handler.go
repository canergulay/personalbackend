package handlers

import "github.canergulay/blogbackend/internal/server/services"

type WebsocketHandler struct {
	createPostService *services.CreatePostService
	savePostService   *services.SavePostService
}

func NewWebSocketHandler(cpsv *services.CreatePostService, svps *services.SavePostService) WebsocketHandler {
	return WebsocketHandler{createPostService: cpsv, savePostService: svps}
}
