package main

import (
	"github.com/lucas11776-golang/calendar_notify/database"
	"github.com/lucas11776-golang/calendar_notify/jobs"
	"github.com/lucas11776-golang/calendar_notify/web"
	"github.com/lucas11776-golang/http/utils/env"
)

func main() {
	env.Load(".env")
	database.Setup()
	go jobs.Run()
	web.Run()
}
