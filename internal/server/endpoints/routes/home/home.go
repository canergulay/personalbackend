package home

import (
	"net/http"

	"github.canergulay/blogbackend/internal/server/endpoints"
	"github.com/labstack/echo/v4"
)

type HomeManager struct {
}

func NewHomeManager() HomeManager {
	return HomeManager{}
}

func (h HomeManager) Handler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func (h HomeManager) GetHomeEndpoint() endpoints.Endpoint {
	return endpoints.NewEndpoint("/", h.Handler, "GET")
}
