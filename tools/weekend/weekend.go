package weekend

import (
	"time"
)

var dateFormat = "2006-01-02"

func getWeekendFromDate(start time.Time) []time.Time {
	var datesAvailability []time.Time
	start.Year()
	end, _ := time.Parse(dateFormat, "2021-12-01")
	end = end.Add(time.Hour * 24)

	for t := start; t.Before(end); t = t.Add(time.Hour * 24) {
		if t.Weekday() == time.Saturday {
			datesAvailability = append(datesAvailability, t)
		}
	}
	return datesAvailability
}
