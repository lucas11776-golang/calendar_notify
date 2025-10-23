package models

type Event struct {
	Connection  string `json:"-" connection:"default" table:"events"`
	ID          string `json:"id" column:"id"`
	StartDate   int    `json:"start_date" column:"start_date"`
	StartEnd    int    `json:"end_date" column:"end_date"`
	Link        string `json:"link" column:"link"`
	Title       string `json:"title" column:"title"`
	Description string `json:"description" column:"description"`
}
