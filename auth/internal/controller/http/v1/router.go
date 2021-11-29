package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(handler *echo.Echo, t interface{}) {
	handler.Use(
		middleware.Logger(),
		middleware.Recover(),
		middleware.RequestID(),
	)
	h := handler.Group("/v1")
	{
		bindRoutes(h, t)
	}
}

func bindRoutes(h *echo.Group, t interface{}) {
	h.GET("/hello", func(ctx echo.Context) error {
		return ctx.JSON(200, map[string]string{"status": "ok"})
	})
	h.POST("/sign-in", func(ctx echo.Context) error {
		return nil
	})
}
