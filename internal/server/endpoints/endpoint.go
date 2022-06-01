package endpoints

import (
	"github.com/labstack/echo/v4"
)

type Endpointable interface {
	GetEndpoint() string
	GetHandler() echo.HandlerFunc
	GetMethod() string
}
