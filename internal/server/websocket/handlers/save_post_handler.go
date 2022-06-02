package handlers

import (
	"github.canergulay/blogbackend/internal/server/routes/blog"
	"github.com/gorilla/websocket"
)

func (wbsckhandler WebsocketHandler) SavePost(data map[string]interface{}, conn *websocket.Conn) {
	post := blog.Post{ID: int(data["id"].(float64)), Content: data["content"].(string), Title: data["title"].(string)}
	wbsckhandler.savePostService.Save(&post)
}
