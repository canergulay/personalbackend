package handlers

import (
	"fmt"

	"github.com/gorilla/websocket"
)

type SavePost struct {
	id    int    "json:id"
	title string "json:title"
	post  string "json:post"
}

func HandleSavePost(conn *websocket.Conn, data interface{}) {
	savePost, ok := data.(SavePost)
	if !ok {
		fmt.Println("data is not in the correct form !")
	}
	fmt.Println(savePost)
}
