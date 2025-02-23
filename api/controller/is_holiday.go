package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"time"

	"github.com/gin-gonic/gin"
)

type APIResponse struct {
	Kind             string        `json:"kind"`
	Etag             string        `json:"etag"`
	Summary          string        `json:"summary"`
	Description      string        `json:"description"`
	Updated          time.Time     `json:"updated"`
	TimeZone         string        `json:"timeZone"`
	AccessRole       string        `json:"accessRole"`
	DefaultReminders []interface{} `json:"defaultReminders"`
	Items            []struct {
		Kind        string    `json:"kind"`
		Etag        string    `json:"etag"`
		ID          string    `json:"id"`
		Status      string    `json:"status"`
		HTMLLink    string    `json:"htmlLink"`
		Created     time.Time `json:"created"`
		Updated     time.Time `json:"updated"`
		Summary     string    `json:"summary"`
		Description string    `json:"description"`
		Creator     struct {
			Email       string `json:"email"`
			DisplayName string `json:"displayName"`
			Self        bool   `json:"self"`
		} `json:"creator"`
		Organizer struct {
			Email       string `json:"email"`
			DisplayName string `json:"displayName"`
			Self        bool   `json:"self"`
		} `json:"organizer"`
		Start struct {
			Date string `json:"date"`
		} `json:"start"`
		End struct {
			Date string `json:"date"`
		} `json:"end"`
		Transparency string `json:"transparency"`
		Visibility   string `json:"visibility"`
		ICalUID      string `json:"iCalUID"`
		Sequence     int    `json:"sequence"`
		EventType    string `json:"eventType"`
	} `json:"items"`
}

// IsHoliday は当日が祝日かどうかを判定するAPI
func IsHoliday(c *gin.Context) {

	timeNow := time.Now()

	url, _ := url.Parse(os.Getenv("HOLIDAY_API_URL"))
	url.Path = path.Join("calendar", "v3", "calendars", os.Getenv("HOLIDAY_GOOGLE_API_CALENDAR_ID"), "events")
	params := url.Query()
	params.Set(
		"key", os.Getenv("HOLIDAY_GOOGLE_API_KEY"),
	)
	params.Set(
		"timeMin", timeNow.Format(time.RFC3339),
	)
	params.Set(
		"timeMax", time.Date(timeNow.Year(), timeNow.Month(), timeNow.Day(), 23, 59, 59, 0, timeNow.Location()).Format(time.RFC3339),
	)
	params.Set(
		"orderBy", "startTime",
	)
	params.Set(
		"singleEvents", "true",
	)
	url.RawQuery = params.Encode()

	// ここでGoogleカレンダーAPIを叩いて祝日かどうかを判定する
	res, err := http.Get(url.String())
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer res.Body.Close()
	data, _ := io.ReadAll(res.Body)

	var response APIResponse
	json.Unmarshal(data, &response)

	for _, item := range response.Items {
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
