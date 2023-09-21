package utils

import (
	"strings"
	"time"
)

type date struct{}

var Date date

const TimeLayout = "2006-01-02T15:04:05.000Z"
const TimeLayout_Default = "2006-01-02 15:04:05.999999999 -0700 MST"

func (d date) ToISOString(t time.Time) string {
	return t.UTC().Format(TimeLayout)
}

func (d date) Parse(s string) (t time.Time, err error) {
	return time.Parse(TimeLayout, s)
}

func (d date) FromIsoString(s string) int64 {
	t, err := time.Parse(TimeLayout, s)
	if err != nil {
		return 0
	}
	return t.Unix()
}
func (d date) ToUnix(i int64) (tm time.Time) {
	tm = time.Unix(i, 0)
	return
}
func (d date) CurrentTimeStampSecond() int64 {
	return time.Now().Unix()
}

func (d date) CurrentDate() string {
	var weekDay = time.Now().UTC().Weekday()
	return strings.ToLower(weekDay.String())
}

func (d date) NextDayAtHour(hour int) int64 {
	now := time.Now()
	tomorrow := time.Date(now.Year(), now.Month(), now.Day()+1, hour, 0, 0, 0, time.UTC)
	return tomorrow.Unix()
}

func (d date) NextWeekAtHour(hour int) int64 {
	now := time.Now()
	iWeekDay := int(now.Weekday())
	remainInWeek := 7 - iWeekDay
	nextWeekAtHour := time.Date(now.Year(), now.Month(), now.Day()+remainInWeek+1, hour, 0, 0, 0, time.UTC)
	return nextWeekAtHour.Unix()
}
