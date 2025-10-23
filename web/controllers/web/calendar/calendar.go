package calendar

import (
	"github.com/lucas11776-golang/calendar_notify/services/calender"
	"github.com/lucas11776-golang/http"
)

// Comment
func Index(req *http.Request, res *http.Response) *http.Response {
	events, err := calender.AllEvents()

	if err != nil {
		return res.SetStatus(http.HTTP_RESPONSE_INTERNAL_SERVER_ERROR)
	}

	return res.View("index", http.ViewData{
		"events": &events,
	})
}
