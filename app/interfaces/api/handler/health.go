package handler

import (
	"time"

	"github.com/tommylike/apaxrag/app/interfaces/api/response"

	"github.com/gin-gonic/gin"
	"github.com/tommylike/apaxrag/app/interfaces/api"
)

type HealthCheck interface {
	Get(c *gin.Context)
}

func NewHealthCheck() HealthCheck {
	return &healthCheck{}
}

type healthCheck struct {
}

func (a *healthCheck) Get(c *gin.Context) {
	api.ResSuccess(c, &response.HealthCheck{
		Status:    "OK",
		CheckedAt: time.Now(),
	})
}
