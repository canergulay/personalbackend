package home

import (
	"net/http"

	"github.canergulay/blogbackend/internal/server/endpoints"
	"github.com/labstack/echo/v4"
)

func handler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func GetHomeEndpoint() endpoints.Endpoint {
	return endpoints.NewEndpoint("/", handler, "GET")
}
