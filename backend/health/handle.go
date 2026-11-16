package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Init(api *gin.RouterGroup) {
	api.GET("/health", Health)
}

func Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
