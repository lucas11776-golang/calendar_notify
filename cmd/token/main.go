package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

// getClient uses a Config to get a Client.
// It checks for a saved token and attempts to refresh if necessary.
func getClient(config *oauth2.Config) *http.Client {
	// 1. Try to read the saved token from a file (e.g., token.json)
	tok, err := tokenFromFile("token.json")

	if err != nil {
		// If no token is found, or it's invalid, get a new one.
		tok = getTokenFromWeb(config)
		saveToken("token.json", tok) // Save the new token for future use
	}

	// 2. Create a TokenSource that automatically refreshes the token.
	tokenSource := config.TokenSource(context.Background(), tok)

	// 3. Return an HTTP client configured with the token source.
	// This client will automatically inject the Access Token and refresh it when needed.
	return oauth2.NewClient(context.Background(), tokenSource)
}

// getTokenFromWeb initiates the Phase 1 browser flow to get a new token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	// Phase 1: Generate the URL and prompt the user to visit it.
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)

	fmt.Printf("Go to the following link in your browser: \n%v\n", authURL)

	// Get the Authorization Code from the user input (after they approve in browser)
	fmt.Printf("Paste the authorization code here: ")

	var authCode string

	if _, err := fmt.Scan(&authCode); err != nil {
		fmt.Printf("Unable to read authorization code: %v", err)
		os.Exit(1)
	}

	// Phase 2: Exchange the code for an Access Token and Refresh Token.
	tok, err := config.Exchange(context.TODO(), authCode)

	if err != nil {
		fmt.Printf("Unable to retrieve token from web: %v", err)
		os.Exit(1)
	}

	return tok
}

// Simple helpers to save and load the token (optional, but essential for production)
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)

	if err != nil {
		return nil, err
	}

	defer f.Close()

	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)

	return tok, err
}

func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		fmt.Printf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

// The scope needed for read-only access to calendar events.
const calendarScope = calendar.CalendarReadonlyScope

func main() {
	// 1. Load your client secrets JSON (downloaded from Google Cloud Console)
	b, err := os.ReadFile("credentials.json")

	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// 2. Configure OAuth2
	config, err := google.ConfigFromJSON(b, calendarScope)

	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}

	// 3. Get the authenticated HTTP client (handles all OAuth steps)
	client := getClient(config)

	// 4. Create the Calendar service client
	srv, err := calendar.NewService(context.Background(), option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Calendar client: %v", err)
	}

	// 5. Build the API call to list events
	// Define a time range for the next 7 days (highly recommended)
	t := time.Now().Format(time.RFC3339)
	sevenDaysLater := time.Now().Add(7 * 24 * time.Hour).Format(time.RFC3339)

	// Use "primary" for the user's main calendar
	events, err := srv.Events.List("primary").
		ShowDeleted(false).
		SingleEvents(true).
		TimeMin(t).
		TimeMax(sevenDaysLater).
		OrderBy("startTime").
		Do()

	if err != nil {
		log.Fatalf("Unable to retrieve next 7 days events: %v", err)
	}

	// 6. Print the events
	if len(events.Items) == 0 {
		fmt.Println("No upcoming events found.")
		return
	}

	fmt.Println("Upcoming events for the next 7 days:")

	for _, item := range events.Items {
		date := item.Start.DateTime
		if date == "" {
			date = item.Start.Date
		}
		fmt.Printf("%s (%s)\n", item.Summary, date)
	}
}
