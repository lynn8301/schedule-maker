package routers

import (
	"github.com/gin-gonic/gin"
	"schedule-maker/controller"
)

func UserRouters() *gin.Engine {
	router := gin.Default()
	userRouter := router.Group("/users")
	{
		userRouter.POST("/", controller.createUser)
	}

	return router
}
