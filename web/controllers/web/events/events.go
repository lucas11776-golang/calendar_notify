package events

import "github.com/lucas11776-golang/http"

// Comment
func Index(req *http.Request, res *http.Response) *http.Response {
	return res.View("events.index", http.ViewData{})
}

// Comment
func View(req *http.Request, res *http.Response) *http.Response {
	return res.View("events.view", http.ViewData{})
}
