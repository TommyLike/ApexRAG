package router

import (
	"github.com/gin-gonic/gin"
	"github.com/tommylike/apaxrag/app/interfaces/api/handler"
)

type Router interface {
	Register(app *gin.Engine) error
	Prefixes() []string
}

func NewRouter(
	userHandler handler.User,
	healthHandler handler.HealthCheck,
) Router {
	return &router{
		userHandler:   userHandler,
		healthHandler: healthHandler,
	}
}

type router struct {
	userHandler   handler.User
	healthHandler handler.HealthCheck
}

func (a *router) Register(app *gin.Engine) error {
	a.RegisterAPI(app)
	return nil
}

func (a *router) Prefixes() []string {
	return []string{
		"/api/",
	}
}
