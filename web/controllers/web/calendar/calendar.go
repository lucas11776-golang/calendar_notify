package calendar

import "github.com/lucas11776-golang/http"

// Comment
func Index(req *http.Request, res *http.Response) *http.Response {

	return res.View("index", http.ViewData{})
}
