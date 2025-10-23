package web

import (
	"fmt"
	"io"
	"os"

	"github.com/lucas11776-golang/calendar_notify/web/controllers/web/calendar"
	"github.com/lucas11776-golang/calendar_notify/web/controllers/web/events"
	"github.com/lucas11776-golang/http"
	"github.com/lucas11776-golang/http/utils/env"
)

func Run() {
	server := http.Server(env.Env("HOST"), env.EnvInt("PORT")).
		SetStatic(env.Env("ASSETS")).
		SetView(env.Env("VIEWS"), "html")

	server.Route().Group("/", func(route *http.Router) {
		route.Get("/", calendar.Index)
		route.Group("events", func(route *http.Router) {
			route.Get("/", events.Index)
		})

		route.Get("auth/google/callback", func(req *http.Request, res *http.Response) *http.Response {
			file, err := os.Create("credentials.json")

			if err != nil {
				fmt.Println("Error open token file", err)
				return res
			}

			data, err := io.ReadAll(req.Body)

			if err != nil {
				fmt.Println("Fail to read body", err)
				return res
			}

			if _, err := file.Write(data); err != nil {
				fmt.Println("Fail to write file", err)
				return res
			}

			return res
		})

		route.Fallback(func(req *http.Request, res *http.Response) *http.Response {
			return res.SetStatus(http.HTTP_RESPONSE_NOT_FOUND).
				View("404", http.ViewData{})
		})
	})

	fmt.Printf("Running Server: %s", server.Host())

	server.Listen()
}
