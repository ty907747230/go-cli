package routes

import (
	"net/http"
	"web_opp/logger"

	"github.com/gin-gonic/gin"
)

//路由
func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello q1mi!",
		})
	})
	return r
}
