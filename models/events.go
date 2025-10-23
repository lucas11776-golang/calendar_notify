package models

import "time"

type Event struct {
	Connection     string `json:"-" connection:"default" table:"events"`
	ID             string `json:"id" column:"id"`
	StartTimestamp int64  `json:"start_timestamp" column:"start_timestamp"`
	EndTimestamp   int64  `json:"end_timestamp" column:"end_timestamp"`
	Link           string `json:"link" column:"link"`
	Title          string `json:"title" column:"title"`
	Description    string `json:"description" column:"description"`
}

// Comment
func (ctx *Event) StartTime() time.Time {
	return time.Unix(ctx.StartTimestamp/1000, (ctx.StartTimestamp%1000)*1_000_000)
}

// Comment
func (ctx *Event) EndTime() time.Time {
	return time.Unix(ctx.EndTimestamp/1000, (ctx.EndTimestamp%1000)*1_000_000)
}
