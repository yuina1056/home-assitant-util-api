package controller

import (
	googlecalendar "home-assitant-util-api/google_calendar"
	"time"

	"github.com/gin-gonic/gin"
)

// IsHoliday は当日が祝日または祝日扱いかどうかを判定するAPI
// 中部電力のスマートライフプランにて休日料金として扱われる日は祝日扱いとする
func IsHolidayRate(c *gin.Context) {

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
				"date":        timeNow.Format(time.DateOnly),
				"holidayRate": true,
			})
			return
		}
	}

	// 祝日でない場合でも特定日の場合は祝日と扱い、trueを返す
	specialDay := []struct {
		Month int
		Day   int
	}{
		// 中部電力のスマートライフプランにて休日料金として扱われる日
		{1, 2},   // 1月2日
		{1, 3},   // 1月3日
		{4, 30},  // 4月30日
		{5, 1},   // 5月1日
		{5, 2},   // 5月2日
		{12, 30}, // 12月30日
		{12, 31}, // 12月31日
	}
	for _, day := range specialDay {
		if timeNow.Month() == time.Month(day.Month) && timeNow.Day() == day.Day {
			c.JSON(200, gin.H{
				"date":        timeNow.Format(time.DateOnly),
				"holidayRate": true,
			})
			return
		}
	}

	c.JSON(200, gin.H{
		"date":        timeNow.Format(time.DateOnly),
		"holidayRate": false,
	})
}
