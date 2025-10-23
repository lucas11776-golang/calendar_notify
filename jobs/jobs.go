package jobs

import (
	"fmt"
	"time"

	"github.com/lucas11776-golang/calendar_notify/services/calender"
)

func Run() {
	// Get calender events...
	go func() {
		for {
			if err := calender.FetchEventsFromGoogleCalender(); err != nil {
				fmt.Println("ERROR FETCHING EVENTS FROM GOOGLE CALENDAR -", err)
				time.Sleep(time.Second * 60)
				continue
			}
			time.Sleep((time.Second * 60) * 30)
		}
	}()

	// Check check event within one hours
	go func() {
		for {

		}
	}()
}
