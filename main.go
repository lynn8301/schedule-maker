package main

import (
	routers "schedule-maker/route"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routers.UserRouters(router)
	routers.ScheduleRouters(router)
	router.Run(":8081")
}
