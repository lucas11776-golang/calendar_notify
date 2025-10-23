package types

type DefaultReminders struct {
	Method  string `json:"method"`
	Minutes int    `json:"minutes"`
}

type Email struct {
	Email string `json:"email"`
}

type Timezone struct {
	DateTime string `json:"dateTime"`
	TimeZone string `json:"timeZone"`
}

type Attendee struct {
	Email          string `json:"email"`
	ResponseStatus string `json:"responseStatus"`
}

type EntryPoint struct {
	EntryPointType string `json:"entryPointType"`
	URI            string `json:"uri"`
	Label          string `json:"label"`
	Pin            string `json:"pin"`
	RegionCode     string `json:"regionCode"`
}

type ConferenceSolutionKey struct {
	Type string `json:"type"`
}

type ConferenceSolution struct {
	Key     ConferenceSolutionKey `json:"key"`
	Name    string                `json:"name"`
	IconUri string                `json:"iconUri"`
}

type Reminders struct {
	UseDefault bool `json:"useDefault"`
}

type ConferenceData struct {
	EntryPoints        []EntryPoint       `json:"entryPoints"`
	ConferenceSolution ConferenceSolution `json:"conferenceSolution"`
	ConferenceId       string             `json:"conferenceId"`
	Reminders          Reminders          `json:"reminders"`
	EventType          string             `json:"eventType"`
}

type Event struct {
	Kind           string         `json:"kind"`
	Etag           string         `json:"etag"`
	ID             string         `json:"id"`
	Status         string         `json:"status"`
	HtmlLink       string         `json:"htmlLink"`
	Created        string         `json:"created"`
	Updated        string         `json:"updated"`
	Summary        string         `json:"summary"`
	Description    string         `json:"description"`
	Creator        Email          `json:"creator"`
	Organizer      Email          `json:"organizer"`
	Start          Timezone       `json:"start"`
	End            Timezone       `json:"end"`
	ICalUID        string         `json:"iCalUID"`
	Sequence       int            `json:"sequence"`
	Attendees      []Attendee     `json:"attendees"`
	HangoutLink    string         `json:"hangoutLink"`
	ConferenceData ConferenceData `json:"conferenceData"`
}

type Events struct {
	Kind             string             `json:"kind"`
	Etag             string             `json:"etag"`
	Summary          string             `json:"summary"`
	Description      string             `json:"description"`
	Updated          string             `json:"updated"`
	Timezone         string             `json:"timeZone"`
	AccessRole       string             `json:"accessRole"`
	DefaultReminders []DefaultReminders `json:"defaultReminders"`
	NextSyncToken    string             `json:"nextSyncToken"`
	Items            []Event            `json:"items"`
}
