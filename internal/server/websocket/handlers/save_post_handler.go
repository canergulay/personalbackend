package handlers

import (
	"encoding/json"

	"github.canergulay/blogbackend/internal/server/routes/blog"
	"github.com/gorilla/websocket"
)

func (wbsckhandler WebsocketHandler) SavePost(data []byte, conn *websocket.Conn) int {
	var post blog.Post
	json.Unmarshal(data, &post)
	return wbsckhandler.savePostService.Save(&post).ID
}
