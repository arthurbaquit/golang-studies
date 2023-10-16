package main

import (
	"fmt"
	"time"
)

func main() {
	payload := map[string]string{
		"requested_at_utc": "2023-06-20T09:00:00+09:00",
	}
	now := time.Now()
	hoursAndMinutesNow := now.Format("15:04")
	fmt.Println(hoursAndMinutesNow)
	fmt.Println(payload["requested_at_utc"])
	newScheduleDate, _ := time.Parse("2006-01-02T15:04:05-07:00", payload["requested_at_utc"])
	fmt.Println(newScheduleDate)

}
