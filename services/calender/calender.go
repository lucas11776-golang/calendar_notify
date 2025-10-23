package calender

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/lucas11776-golang/calendar_notify/models"
	"github.com/lucas11776-golang/calendar_notify/types"
	"github.com/lucas11776-golang/calendar_notify/utils/token"
	"github.com/lucas11776-golang/orm"
)

const GOOGLE_CALENDAR_EVENTS_URL = "https://www.googleapis.com/calendar/v3/calendars/primary/events"

type EventsFilter struct {
	Search string
	Page   int64
	Limit  int64
}

// comment
func GetEvents(filter EventsFilter) (*orm.Pagination[*models.Event], error) {
	// TODO: must use timezone
	// current, err := time.Parse("2006-01-02T15:04:05-07:00", "2025-10-21T11:30:00+02:00")
	current, err := time.Parse("2006-01-02T15:04:05-07:00", "2025-01-21T11:30:00+02:00")

	fmt.Println("ERROR", err, current.UnixMilli())

	events, err := orm.Model(models.Event{}).
		// Where("end_timestamp", ">", time.Now().UnixMilli()).
		Where("end_timestamp", ">", current.UnixMilli()).
		OrderBy("start_timestamp", "ASC").
		Limit(filter.Limit).
		Paginate(12, filter.Page)

	if err != nil {
		return nil, err
	}

	return events, nil
}

// comment
func AllEvents() ([]*models.Event, error) {
	// TODO: must use timezone
	current, err := time.Parse("2006-01-02T15:04:05-07:00", "2025-10-21T11:30:00+02:00")

	fmt.Println("ERROR", err, current.UnixMilli())

	events, err := orm.Model(models.Event{}).
		// Where("end_timestamp", ">", time.Now().UnixMilli()).
		Where("end_timestamp", ">", current.UnixMilli()).
		OrderBy("start_timestamp", "ASC").
		Get()

	if err != nil {
		return nil, err
	}

	return events, nil
}

// Comment
func CreateOrUpdateEvent(event types.Event) error { // TODO: maybe add FirstAndUpdate -> which check if give value match if not update
	start, err := time.Parse("2006-01-02T15:04:05-07:00", event.Start.DateTime)

	if err != nil {
		return err
	}

	end, err := time.Parse("2006-01-02T15:04:05-07:00", event.End.DateTime)

	entity, err := orm.Model(models.Event{}).
		Where("id", "=", event.ID).
		First()

	if err != nil {
		return err
	}

	if entity == nil {
		// TODO: must use timezone
		entity, err = orm.Model(models.Event{}).
			Insert(orm.Values{
				"id":              event.ID,
				"start_timestamp": start.UnixMilli(),
				"end_timestamp":   end.UnixMilli(),
				"link":            event.HangoutLink,
				"title":           event.Summary,
				"Description":     regexp.MustCompile(`<[^>]*>`).ReplaceAllString(event.Description, ""),
				"attended":        false,
			})

		if err != nil {
			return err
		}
	}

	if entity.StartTimestamp != start.UnixMilli() || entity.EndTimestamp != end.UnixMilli() {
		// TODO: must use timezone
		err := orm.Model(models.Event{}).
			Where("id", "=", event.ID).
			Update(orm.Values{
				"start_timestamp": start.UnixMilli(),
				"end_timestamp":   end.UnixMilli(),
				"link":            event.HangoutLink,
				"title":           event.Summary,
				"Description":     regexp.MustCompile(`<[^>]*>`).ReplaceAllString(event.Description, ""),
				"attended":        false,
			})

		if err != nil {
			return err
		}
	}

	return nil
}

// comment
func FetchEventsFromGoogleCalender() error {
	token, err := token.Get()

	if err != nil {
		return err
	}

	payload := strings.NewReader("")
	client := &http.Client{}
	req, err := http.NewRequest("GET", GOOGLE_CALENDAR_EVENTS_URL, payload)

	if err != nil {
		return err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))
	res, err := client.Do(req)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return err
	}

	var events types.Events

	if err := json.Unmarshal(body, &events); err != nil {
		return err
	}

	for _, event := range events.Items {
		if err := CreateOrUpdateEvent(event); err != nil {
			fmt.Println("FAILED TO UPDATE OR CREATE EVENT:", event.ID, err)
		}
	}

	return nil
}
