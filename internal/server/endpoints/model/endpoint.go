package endpoints

import (
	"github.com/labstack/echo/v4"
)

type Endpoint interface {
	GetEndpoint() string
	GetHandler() echo.HandlerFunc
	GetMethod() string
}
