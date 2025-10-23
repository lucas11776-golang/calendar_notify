package models

type Token struct {
	Connection  string `json:"-" connection:"default" table:"tokens"`
	Name        string `json:"name" column:"name"`
	Expires     int    `json:"expires" column:"expires"`
	AccessToken string `json:"access_token" column:"access_token"`
}
