package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func status(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "OK!"})
}

func InitRouter(repo *RouterHandler) *gin.Engine {
	r := gin.Default()
	r.GET("/", status)
	return r
}
