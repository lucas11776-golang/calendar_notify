package jobs

import (
	"fmt"
	"time"

	"github.com/lucas11776-golang/calendar_notify/models"
	"github.com/lucas11776-golang/calendar_notify/services/calender"
	"github.com/lucas11776-golang/calendar_notify/utils/notification"
	"github.com/lucas11776-golang/http/utils/env"
	"github.com/lucas11776-golang/orm"
)

type Period string

const (
	FUTURE      Period = "future"
	UPCOMING    Period = "up_coming"
	IN_PROGRESS Period = "in_progress"
	END         Period = "end"
)

// Comment
func eventNotifyLoop(event *models.Event) {
	for {
		switch getEventPeriod(event) {
		case UPCOMING:
			notification.New(
				fmt.Sprintf("%s - ", event.Title),
				fmt.Sprintf("Hi, you have upcoming meeting in %d minutes please go to %s", 10, env.Env("APP_URL")),
				"",
			).Show()

		case IN_PROGRESS:
			notification.New(
				fmt.Sprintf("%s", event.Title),
				fmt.Sprintf("Hi, the is a meeting in progress please go to %s", env.Env("APP_URL")),
				"",
			).Show()

		default:
			delete(queue, event.ID)
		}

		// Meeting may be changed...
		if ev, err := orm.Model(models.Event{}).Where("id", "=", event.ID).First(); err == nil {
			event = ev
		}

		time.Sleep(time.Minute * 1)
	}

}

// Comment
func getEventPeriod(event *models.Event) Period {
	current_time := time.Now().UnixMilli()

	if current_time >= event.EndTimestamp {
		return END
	}

	if event.StartTimestamp >= current_time || current_time <= event.EndTimestamp {
		return IN_PROGRESS
	}

	timeToMeeting := getTimeToMeeting(event)

	if timeToMeeting > 0 && timeToMeeting <= (time.Minute*30).Milliseconds() {
		return UPCOMING
	}

	return FUTURE
}

// Comment
func getTimeToMeeting(event *models.Event) int64 {
	return event.StartTimestamp - time.Now().UnixMilli()
}

var queue map[string]*models.Event = make(map[string]*models.Event)

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
			events, err := calender.AllEvents()

			if err != nil {
				notification.New("Failed to get events", err.Error(), "").Show()
			}

			for _, event := range events {
				if _, in := queue[event.ID]; in {
					continue
				}

				switch getEventPeriod(event) {
				case UPCOMING, IN_PROGRESS:
					queue[event.ID] = event

					fmt.Println("IN", event.Title, event.ID)
					eventNotifyLoop(event)
				}
			}

			time.Sleep(time.Minute * 5)
		}
	}()
}
