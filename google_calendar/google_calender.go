package googlecalendar

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"time"
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

func GetHoliday(timeMin time.Time, timeMax time.Time) (*APIResponse, error) {
	url, _ := url.Parse(os.Getenv("HOLIDAY_API_URL"))
	url.Path = path.Join("calendar", "v3", "calendars", os.Getenv("HOLIDAY_GOOGLE_API_CALENDAR_ID"), "events")
	params := url.Query()
	params.Set(
		"key", os.Getenv("HOLIDAY_GOOGLE_API_KEY"),
	)
	params.Set(
		"timeMin", timeMin.Format(time.RFC3339),
	)
	params.Set(
		"timeMax", timeMax.Format(time.RFC3339),
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
		return nil, err
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response APIResponse
	if err := json.Unmarshal(data, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
