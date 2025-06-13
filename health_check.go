package server

import (
	"net/http"
	"sync/atomic"

	"github.com/gin-gonic/gin"
)

type HealthCheck struct {
	IsUnhealthy *atomic.Bool
}

func NewHealthCheck() *HealthCheck {
	return &HealthCheck{
		IsUnhealthy: &atomic.Bool{},
	}
}

func (h *HealthCheck) Health(c *gin.Context) {
	if !h.IsUnhealthy.Load() {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
		return
	}
	c.JSON(http.StatusInternalServerError, gin.H{"status": "unhealthy"})
}

func (h *HealthCheck) SetUnhealthy() {
	h.IsUnhealthy.Store(true)
}
