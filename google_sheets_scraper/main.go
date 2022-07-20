package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

func main() {
	ctx := context.Background()
	b, err := ioutil.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, "https://www.googleapis.com/auth/spreadsheets.readonly")
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := getClient(config)

	srv, err := sheets.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	spreadsheetId := "10-eMLMp7EGFdQd9rZW1Bb5TzB7AN00uDvg__40_M05c"
	readRange := "Pelicans!A2:E"
	resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	cards := scrapeCards(resp, "Pelicans")
	readRange = "LSU!A2:E"
	resp, err = srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	cards2 := scrapeCards(resp, "LSU")

	seedToMongo(cards, "Pelicans")
	seedToMongo(cards2, "LSU")
	fmt.Println("Seeding done.")
}
