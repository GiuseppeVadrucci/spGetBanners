package api

import (
	"handlers"

	"github.com/labstack/echo"
)

func MainGroup(e *echo.Echo) {
	// Route / to handler function
	e.GET("/health-check", handlers.HealthCheck)

	e.GET("/cms/banners:data", handlers.GetBann)
	e.POST("/cms/banners", handlers.AddBann)

}

func AdminGroup(g *echo.Group) {
	g.GET("/main", handlers.MainAdmin)
}
