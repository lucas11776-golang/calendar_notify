package token

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/lucas11776-golang/calendar_notify/models"
	"github.com/lucas11776-golang/calendar_notify/types"
	"github.com/lucas11776-golang/http/utils/env"
	"github.com/lucas11776-golang/orm"
)

// Comment
func Refresh() (*types.Token, error) {
	jsonData, err := json.Marshal(map[string]string{
		"client_id":     env.Env("GOOGLE_CLIENT_ID"),
		"client_secret": env.Env("GOOGLE_CLIENT_SECRET"),
		"refresh_token": env.Env("GOOGLE_REFRESH_TOKEN"),
		"grant_type":    "refresh_token",
	})

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", env.Env("GOOGLE_REFRESH_TOKEN_URL"), bytes.NewBuffer(jsonData))

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var token types.Token

	if err := json.Unmarshal(body, &token); err != nil {
		return nil, err
	}

	return &token, nil
}

// comment
func UpdateOrCreate(token *types.Token) (*models.Token, error) { // TODO: orm should have UpdateOrCreate
	entity, err := orm.Model(models.Token{}).
		Where("name", "=", env.Env("GOOGLE_TOKEN_CALENDAR_NAME")).
		First()

	if err != nil {
		return nil, err
	}

	if entity == nil {
		return orm.Model(models.Token{}).
			Insert(orm.Values{
				"name":         env.Env("GOOGLE_TOKEN_CALENDAR_NAME"),
				"access_token": token.AccessToken,
				"expires":      ((token.ExpiresIn - 60) * 1000) * int(time.Now().UnixMilli()),
			})
	}

	err = orm.Model(models.Token{}).
		Where("name", "=", env.Env("GOOGLE_TOKEN_CALENDAR_NAME")).
		Update(orm.Values{
			"access_token": token.AccessToken,
			"expires":      ((token.ExpiresIn - 60) * 1000) * int(time.Now().UnixMilli()),
		})

	if err != nil {
		return nil, err
	}

	entity.AccessToken = token.AccessToken
	entity.Expires = ((token.ExpiresIn - 60) * 1000) * int(time.Now().UnixMilli())

	return entity, nil
}

// Comment
func Create() (*models.Token, error) {
	token, err := Refresh()

	if err != nil {
		return nil, err
	}

	return UpdateOrCreate(token)
}

// Comment
func Get() (*models.Token, error) {
	entity, err := orm.Model(models.Token{}).
		Where("name", "=", env.Env("GOOGLE_TOKEN_CALENDAR_NAME")).
		First()

	if err != nil {
		return nil, err
	}

	if entity == nil || entity.Expires < int(time.Now().UnixMilli()) {
		return Create()
	}

	return entity, nil
}
