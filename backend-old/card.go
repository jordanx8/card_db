package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"google.golang.org/api/sheets/v4"
)

type Card struct {
	FirstName     string `json:"FirstName"`
	LastName      string `json:"LastName"`
	SeasonsPlayed string `json:"SeasonsPlayed"`
	Season        string `json:"Season"`
	Manufacturer  string `json:"Manufacturer"`
	Set           string `json:"Set"`
	Insert        string `json:"Insert"`
	Parallel      string `json:"Parallel"`
	CardNumber    string `json:"CardNumber"`
	Notes         string `json:"Notes"`
}

func scrapeCards(resp *sheets.ValueRange, sheetName string) []Card {
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
					split := strings.Fields(row[0].(string))
					c.FirstName = split[0]
					c.LastName = split[1]
					c.SeasonsPlayed = split[2]
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
				c.Notes = ""
				if len(row) == 5 {
					c.Notes = row[4].(string)
					// notesArr := strings.FieldsFunc(row[4].(string), func(r rune) bool {
					// 	if r == ',' {
					// 		return true
					// 	}
					// 	return false
					// })
					// for _, note := range notesArr {
					// 	c.Notes = append(c.Notes, strings.TrimSpace(note))
					// }
				}
				ready = true
			}
			if ready {
				cards = append(cards, c)
				ready = false
			}
		}
		file, _ := json.MarshalIndent(cards, "", " ")
		_ = ioutil.WriteFile("../frontend/src/"+sheetName+".json", file, 0644)
		fmt.Println(len(cards), "cards were scraped.")
	}
	return cards
}
