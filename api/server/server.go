package server

import (
	"home-assitant-util-api/api/controller"

	"github.com/gin-gonic/gin"
)

func Handler(port string) {
	r := gin.Default()

	r.GET("/isholiday", controller.IsHoliday)

	r.Run(":" + port)
}
