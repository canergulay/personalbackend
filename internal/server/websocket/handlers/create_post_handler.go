package handlers

import "github.com/gorilla/websocket"

func (wbsckhandler WebsocketHandler) CreatePost(conn *websocket.Conn) {
	postCreated := wbsckhandler.createPostService.Create()
	response := WSResponse{Status: 0, Data: postCreated}
	conn.WriteJSON(response)
}
