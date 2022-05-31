package server

import (
	"fmt"
	"net/http"

	endpoints "github.canergulay/blogbackend/internal/server/endpoints"
	"github.com/labstack/echo/v4"
)

type ServerManager struct {
	e echo.Echo
}

func GetNewServerManager(endpoints ...endpoints.Endpoint) ServerManager {
	e := echo.New()
	for _, endpoint := range endpoints {

		fmt.Println(endpoint.GetMethod())
		switch endpoint.GetMethod() {
		case http.MethodGet:
			e.GET(endpoint.GetEndpoint(), endpoint.GetHandler())
		case http.MethodPost:
			e.POST(endpoint.GetEndpoint(), endpoint.GetHandler())
		}
	}
	return ServerManager{e: *e}
}

func (server ServerManager) StartServer(urlAndPort string) {
	server.e.Logger.Fatal(server.e.Start(urlAndPort))
}
