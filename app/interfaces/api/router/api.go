package router

import (
	"github.com/gin-gonic/gin"
	"github.com/tommylike/apaxrag/app/interfaces/api/middleware"
)

// RegisterAPI register api group router
func (a *router) RegisterAPI(app *gin.Engine) {
	g := app.Group("/api")
	g.Use(middleware.RateLimiterMiddleware())
	g.GET("health", a.healthHandler.Get)
	v1 := g.Group("/v1")
	{
		gUser := v1.Group("users")
		{
			gUser.GET("", a.userHandler.Query)
			gUser.GET(":id", a.userHandler.Get)
			gUser.POST("", a.userHandler.Create)
			gUser.PUT(":id", a.userHandler.Update)
			gUser.DELETE(":id", a.userHandler.Delete)
			gUser.PATCH(":id/enable", a.userHandler.Enable)
			gUser.PATCH(":id/disable", a.userHandler.Disable)
		}
	}
}
