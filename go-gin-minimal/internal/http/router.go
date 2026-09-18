package httpapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"service": "go-gin-minimal",
		})
	})

	router.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	return router
}
