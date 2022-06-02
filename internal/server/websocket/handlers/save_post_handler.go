package handlers

import (
	"github.canergulay/blogbackend/internal/server/routes/blog"
	"github.com/gorilla/websocket"
)

func (wbsckhandler WebsocketHandler) SavePost(data interface{}, conn *websocket.Conn) {
	dataParsed := data.(map[string]interface{})
	post := blog.Post{ID: int(dataParsed["id"].(float64)), Content: dataParsed["content"].(string), Title: dataParsed["title"].(string)}
	wbsckhandler.savePostService.Save(&post)
}
