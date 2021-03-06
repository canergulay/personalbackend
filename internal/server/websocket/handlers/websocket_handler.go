package handlers

import "github.canergulay/blogbackend/internal/server/services"

type WSResponse struct {
	Status int         `json:"status"` // 0 -> SUCCESS , 1 -> SOMETHING UNEXPECTED HAS HAPPENNED...
	Data   interface{} `json:"data"`
}

type WebsocketHandler struct {
	createPostService *services.CreatePostService
	savePostService   *services.SavePostService
	getPostsService   *services.GetPostsService
}

func NewWebSocketHandler(cpsv *services.CreatePostService, svps *services.SavePostService, gpsv *services.GetPostsService) WebsocketHandler {
	return WebsocketHandler{createPostService: cpsv, savePostService: svps}
}
