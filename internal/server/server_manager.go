package server

import (
	"net/http"

	endpoints "github.canergulay/blogbackend/internal/server/endpoints"
	"github.canergulay/blogbackend/internal/server/endpoints/routes/home"
	"github.canergulay/blogbackend/internal/server/endpoints/routes/websocket"
	"github.com/labstack/echo/v4"
)

type ServerManager struct {
	e echo.Echo
}

func InitialiseAllRoutes() ServerManager {
	ws := websocket.GetWebsocketHandler()
	home := home.GetHomeEndpoint()
	return getNewServerManager(home, ws)
}

func getNewServerManager(endpoints ...endpoints.Endpoint) ServerManager {
	e := echo.New()

	// INJECTION OF ALL THE ENDPOINTS
	for _, endpoint := range endpoints {
		switch endpoint.GetMethod() {
		case http.MethodGet:
			e.GET(endpoint.GetEndpoint(), endpoint.GetHandler())
		case http.MethodPost:
			e.POST(endpoint.GetEndpoint(), endpoint.GetHandler())
		}
	}
	// SERVERMANAGER WHICH CONTAINS THE ECHO WHICH ALL ENDPOINTS ARE INJECTED IN !
	return ServerManager{e: *e}
}

func (server ServerManager) StartServer(urlAndPort string) {
	server.e.Logger.Fatal(server.e.Start(urlAndPort))
}
