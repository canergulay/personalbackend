package websocket

import (
	"fmt"
	"net/http"

	"github.canergulay/blogbackend/internal/server/endpoints"
	"github.canergulay/blogbackend/internal/server/endpoints/routes/websocket/handlers"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func handler(c echo.Context) error {
	// TO PREVENT CORS ERRORS
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	// THIS IS ONE OF THE BENEFITS OF USING GO
	// I WAS USING WEBSOCKETS IN NODEJS AND DART BEFORE BUT
	// I DIDNT'T KNOW THE ESSENCE OF IT & HOW IT WORKS IN THE BACK, WHAT'S BEHIND IT
	// I ALWAYS THOUGHT IT IS TOTALLY A SEPARATE PROTOCOL
	// BUT NOW , I KNOW THAT IT IS NOTHING BUT A UPGRADED HTTP GET REQUEST
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)

	if err != nil {
		fmt.Println(err)

		return err
	}

	defer ws.Close()

	for {
		_, msg, err := ws.ReadMessage()

		if err != nil {
			return err
		}

		err2, parsedMessage := ParseMessage(msg)

		if err2 != nil {
			return err
		}

		switch parsedMessage.Type {
		case CreatePost:
			handlers.CreatePostHandler(ws)
		case SavePost:
			handlers.SavePostHandler(ws, parsedMessage)
		}
	}

}

func GetWebsocketHandler() endpoints.Endpoint {
	return endpoints.NewEndpoint("/ws", handler, "GET")
}
