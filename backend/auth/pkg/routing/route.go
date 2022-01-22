package routing

import (
	"github.com/labstack/echo/v4"
)

func BindHttpRouteGroup(app *echo.Echo, groups ...IHttpRouteGroup) {
	for _, group := range groups {
		endpointGroup := app.Group(group.Prefix(), group.Middleware()...)
		{
			for _, handler := range group.Handlers() {
				endpointGroup.Add(
					handler.Method(),
					handler.Endpoint(),
					handler.Handler(),
					handler.Middleware()...,
				)
			}
		}
	}
}
