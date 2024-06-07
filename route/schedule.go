package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getSchdeule(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"msg":    "FROM schedule",
	})
}

func ScheduleRouters(router *gin.Engine) {
	ScheduleRouters := router.Group("/schedule")
	{
		ScheduleRouters.GET("", getSchdeule)
	}
}
