package server

import (
	"net/http"

	"github.canergulay/blogbackend/internal/server/endpoints"
	"github.canergulay/blogbackend/internal/server/routes/blog"
	"github.canergulay/blogbackend/internal/server/routes/home"
	"github.canergulay/blogbackend/internal/server/websocket"
	"github.com/labstack/echo/v4"
)

type ServerManager struct {
	e echo.Echo
}

func InitialiseAllRoutes(
	blogManager *blog.BlogManager,
	socketManager *websocket.SocketManager,
	homeManager *home.HomeManager) ServerManager {

	return getNewServerManager(
		homeManager.GetHomeEndpoint(),
		socketManager.GetWebsocketService(),
		blogManager.GetBlogEndpoint())
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
