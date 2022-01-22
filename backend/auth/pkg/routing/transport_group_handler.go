package routing

import "github.com/labstack/echo/v4"

type HttpRouteHandler struct {
	endpoint   string
	method     string
	handler    echo.HandlerFunc
	middleware []echo.MiddlewareFunc
}

func NewHttpRouteHandler(endpoint string, method string, handler echo.HandlerFunc, middleware ...echo.MiddlewareFunc) *HttpRouteHandler {
	return &HttpRouteHandler{endpoint: endpoint, method: method, handler: handler, middleware: middleware}
}

func (t *HttpRouteHandler) Endpoint() string {
	return t.endpoint
}

func (t *HttpRouteHandler) Method() string {
	return t.method
}

func (t *HttpRouteHandler) Handler() echo.HandlerFunc {
	return t.handler
}

func (t *HttpRouteHandler) Middleware() []echo.MiddlewareFunc {
	return t.middleware
}
