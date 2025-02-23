package controller

import (
	googlecalendar "home-assitant-util-api/google_calendar"
	"time"

	"github.com/gin-gonic/gin"
)

// IsHoliday は当日が祝日かどうかを判定するAPI
func IsHoliday(c *gin.Context) {

	timeNow := time.Now()

	response, err := googlecalendar.GetHoliday(
		timeNow,
		time.Date(timeNow.Year(), timeNow.Month(), timeNow.Day(), 23, 59, 59, 0, timeNow.Location()),
	)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	for _, item := range response.Items {
		// 祝日の場合はtrueを返す
		if item.Start.Date == timeNow.Format(time.DateOnly) {
			c.JSON(200, gin.H{
				"date":    timeNow.Format(time.DateOnly),
				"holiday": true,
			})
			return
		}
	}

	c.JSON(200, gin.H{
		"date":    timeNow.Format(time.DateOnly),
		"holiday": false,
	})
}
