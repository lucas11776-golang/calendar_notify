package events

import (
	"github.com/lucas11776-golang/calendar_notify/services/calender"
	"github.com/lucas11776-golang/http"
	"github.com/spf13/cast"
)

// Comment
func Index(req *http.Request, res *http.Response) *http.Response {
	events, err := calender.GetEvents(calender.EventsFilter{
		Search: "",
		Page:   cast.ToInt64(req.GetQuery("page")),
	})

	if err != nil {
		return res.SetStatus(http.HTTP_RESPONSE_INTERNAL_SERVER_ERROR)
	}

	return res.View("events.index", http.ViewData{
		"events": &events,
	})
}
