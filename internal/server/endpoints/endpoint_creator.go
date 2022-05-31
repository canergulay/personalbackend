package endpoints

import "github.com/labstack/echo/v4"

type Endpoint struct {
	Endpoint string
	Handler  echo.HandlerFunc
	Method   string
}

func NewEndpoint(endpoint string, handler echo.HandlerFunc, method string) Endpoint {
	return Endpoint{Endpoint: endpoint, Handler: handler, Method: method}
}

func (endpnt Endpoint) GetEndpoint() string {
	return endpnt.Endpoint
}

func (endpnt Endpoint) GetHandler() echo.HandlerFunc {
	return endpnt.Handler
}

func (endpnt Endpoint) GetMethod() string {
	return endpnt.Method
}
