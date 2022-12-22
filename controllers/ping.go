package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PingSwag struct {
	Health string `example:"ok" json:"health"`
}

// @BasePath /api/v1
// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags Health
// @Accept json
// @Produce json
// @Success 200 {object} PingSwag
// @Router /ping [get]
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"health": "OK",
	})
}
