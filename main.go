package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

type Card struct {
	PlayerName   string   `json:"Player Name"`
	Season       string   `json:"Season"`
	Manufacturer string   `json:"Manufacturer"`
	Set          string   `json:"Set"`
	Insert       string   `json:"Insert"`
	Parallel     string   `json:"Parallel"`
	CardNumber   string   `json:"Card Number"`
	Notes        []string `json:"Notes"`
}

// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config) *http.Client {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tokFile := "token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

// Retrieves a token from a local file.
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

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

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

	// Prints the names and majors of students in a sample spreadsheet:
	// https://docs.google.com/spreadsheets/d/1BxiMVs0XRA5nFMdKvBdBZjgmUUqptlbs74OgvE2upms/edit
	spreadsheetId := "10-eMLMp7EGFdQd9rZW1Bb5TzB7AN00uDvg__40_M05c"
	readRange := "Pelicans!A2:E"
	resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()

	var c Card
	var cards []Card
	var ready = false
	if len(resp.Values) == 0 {
		fmt.Println("No data found.")
	} else {
		for _, row := range resp.Values {
			if len(row) == 1 {
				_, err := strconv.Atoi(row[0].(string)[0:4])
				if err != nil {
					//name
					// fmt.Printf("Player: %s\n", row)
					c.PlayerName = row[0].(string)
				} else {
					//year
					// fmt.Printf("Season: %s\n", row)
					c.Season = row[0].(string)
				}
			} else {
				// fmt.Printf("Card: %s\n", row)
				c.Manufacturer = row[0].(string)
				setinsert := strings.FieldsFunc(row[1].(string), func(r rune) bool {
					if r == '/' {
						return true
					}
					return false
				})
				c.Set = strings.TrimSpace(setinsert[0])
				c.Insert = ""
				if len(setinsert) == 2 {
					c.Insert = strings.TrimSpace(setinsert[1])
				}
				c.Parallel = row[2].(string)
				c.CardNumber = row[3].(string)
				c.Notes = nil
				if len(row) == 5 {
					notesArr := strings.FieldsFunc(row[4].(string), func(r rune) bool {
						if r == ',' {
							return true
						}
						return false
					})
					for _, note := range notesArr {
						c.Notes = append(c.Notes, strings.TrimSpace(note))
					}
				}
				ready = true
			}
			if ready {
				cards = append(cards, c)
				ready = false
			}
		}
		file, _ := json.MarshalIndent(cards, "", " ")
		_ = ioutil.WriteFile("cards.json", file, 0644)
		fmt.Println(len(cards))
	}

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	// Connect to the MongoDB and return Client instance
	mongoClient, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalf("mongo.Connect() ERROR: %v", err)
	}
	mongoCtx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	col := mongoClient.Database("Cards").Collection("Pelicans")
	if err = col.Drop(ctx); err != nil {
		log.Fatal(err)
	}
	col = mongoClient.Database("Cards").Collection("Pelicans") //check for error again
	if err != nil {
		log.Fatal(err)
	} //else mongo has been connected
	fmt.Println("Successfully connected to Mongo")
	for i := range cards {
		card := cards[i]
		fmt.Println("inserting", card)
		// Call the InsertOne() method and pass the context and doc objects
		col.InsertOne(mongoCtx, card)
	}

}
