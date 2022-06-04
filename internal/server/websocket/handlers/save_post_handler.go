package handlers

import (
	"encoding/json"
	"fmt"

	"github.canergulay/blogbackend/internal/server/models"
	"github.com/gorilla/websocket"
)

func (wbsckhandler WebsocketHandler) SavePost(data interface{}, conn *websocket.Conn) {
	dataParsed := data.(map[string]interface{})
	post := models.Post{ID: int(dataParsed["id"].(float64)), Content: dataParsed["content"].(string), Title: dataParsed["title"].(string)}
	wbsckhandler.savePostService.Save(&post)

	// THERE SHOULD DEFINETELY BE ERROR HANDLING HERE
	// BUT WE HAVE TO CONTINUE NOW.
	// THERE IS A TIME AND A PLACE FOR EVERYTHING
	bytes, err := json.Marshal(WSResponse{Status: 0})
	fmt.Println(err)
	fmt.Println("err")

	if err == nil {
		conn.WriteMessage(1, bytes)
	} else {
		fmt.Println(err)
		// TODO : ADD LOGGING.
	}
}
