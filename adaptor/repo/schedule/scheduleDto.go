package schedule

import (
	"time"
	"strings"
)

type (
	Response struct {
		Error bool `json:"error"`
		Data  Data `json:"data"`
	}

	Data struct {
		Schedules []*Schedule `json:"schedules"`
	}

	TimeDto struct {
		time.Time
	}

	Schedule struct {
		Title       string      `json:"title"`
		AllDay      bool        `json:"allDay"`
		Description string      `json:"description"`
		Start       TimeDto     `json:"start"`
		End         TimeDto     `json:"end"`
		ScheduleKey ScheduleKey `json:"scheduleKey"`
	}

	ScheduleKey struct {
		Code string `json:"code"`
	}
)

func (d *TimeDto) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("2006-01-02T15:04:05", s)
	if err != nil {
		return err
	}
	duration, _ := time.ParseDuration("-9h")
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	d.Time = t.Add(duration).In(jst)
	return nil
}
