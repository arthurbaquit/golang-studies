package main

import (
	"fmt"
	"time"
)

var lastMsgSentDay = "2023-05-25"

const (
	INITIAL_WORKING_HOUR = 10
	FINAL_WORKING_HOUR   = 17
)

func main() {
	// For loop for testing purposes (sent once then skip msg case)
	for {
		// now := time.Now()
		brLocation, _ := time.LoadLocation("America/Sao_Paulo")
		now := time.Now().In(brLocation)
		// now := time.Date(2023, 05, 26, 18, 21, 46, 982397, time.UTC)
		if now.Weekday() >= time.Monday && now.Weekday() <= time.Friday &&
			now.Hour() >= INITIAL_WORKING_HOUR && now.Hour() < FINAL_WORKING_HOUR {
			fmt.Println("Skipping any message that isn't out of work time")
			return
		}
		if lastMsgSentDay != "" {
			lastMsgSentDay, err := time.Parse("2006-01-02", lastMsgSentDay)
			fmt.Println(lastMsgSentDay)
			fmt.Println(now)
			if err != nil {
				fmt.Println("Error parsing last message sent day at channel out of work")
				return
			}
			if lastMsgSentDay.Local().Day() == now.Local().Day() {
				fmt.Println("Skipping message because it was sent today")
				return
			}
		}

		msg := `Adicione a msg de out of work aqui`
		fmt.Println(msg)
		lastMsgSentDay = now.Format("2006-01-02")
		fmt.Println(lastMsgSentDay)
	}
}
