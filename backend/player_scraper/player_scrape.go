package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	grpc "github.com/jordanx8/card_db/grpc_server"
	m "github.com/jordanx8/card_db/mongodb"
	"go.mongodb.org/mongo-driver/bson"
)

type PlayerScrape struct {
	FirstName string
	LastName  string
	URL       string
}

var playerInfo []grpc.Player

func getPlayerURLs() []PlayerScrape {
	var url []PlayerScrape

	c := colly.NewCollector(
		colly.AllowedDomains("basketball-reference.com", "www.basketball-reference.com"),
	)

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong: ", err)
	})

	c.OnHTML("tr", func(e *colly.HTMLElement) {
		playerURL := e.ChildAttr("a", "href")
		playerName := e.ChildAttr("td.left", "csk")
		if playerURL != "" {
			names := strings.Split(playerName, ",")
			url = append(url, PlayerScrape{
				FirstName: names[1],
				LastName:  names[0],
				URL:       playerURL,
			})
		}
	})

	err := c.Visit("https://www.basketball-reference.com/teams/NOH/players.html")
	if err != nil {
		log.Printf("failed to visit url: %v\n", err)
	}

	return url
}

func scrapePlayerPage(p PlayerScrape) {
	seasonArray := make([]string, 0)

	c := colly.NewCollector(
		colly.AllowedDomains("basketball-reference.com", "www.basketball-reference.com"),
	)
	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong: ", err)
	})

	c.OnHTML("table#per_game > tbody", func(e *colly.HTMLElement) {
		e.DOM.Find("tr").Each(func(_ int, s *goquery.Selection) {
			season := s.Find("th").Text()
			team := s.Find("td[data-stat='team_id']").Text()
			if season != "" && team != "" && team != "TOT" {
				if (len(seasonArray) != 0) && (strings.Contains(seasonArray[len(seasonArray)-1], season)) {
					prevSeason := seasonArray[len(seasonArray)-1]
					prevSeason = strings.Trim(prevSeason, ")") + ", " + team + ")"
					seasonArray[len(seasonArray)-1] = prevSeason
				} else {
					s := season + " (" + team + ")"
					seasonArray = append(seasonArray, s)
				}
			}
		})
	})

	err := c.Visit("https://www.basketball-reference.com" + p.URL)
	if err != nil {
		log.Printf("failed to visit url: %v\n", err)
	}

	nolaTeams := [3]string{"NOP", "NOH", "NOK"}

	// generate SeasonsPlayed for NOLA
	seasonsPlayed := ""
	index := [2]int{5, 7}
	for _, season := range seasonArray {
		for _, team := range nolaTeams {
			if (strings.Contains(season, team)) && (seasonsPlayed == "") {
				seasonsPlayed = season[:7]
			} else if strings.Contains(season, team) {
				checkEnd, err := strconv.Atoi(seasonsPlayed[index[0]:index[1]])
				if err != nil {
					// ... handle error
					panic(err)
				}

				checkNew, err := strconv.Atoi(season[5:7])
				if err != nil {
					// ... handle error
					panic(err)
				}

				if checkEnd+1 == checkNew {
					seasonsPlayed = seasonsPlayed[:index[0]] + season[5:7]
				} else {
					seasonsPlayed = seasonsPlayed + ", " + season[:7]
					index[0] = index[0] + 9
					index[1] = index[1] + 9
				}
				if strconv.Itoa(checkNew) == "24" {
					seasonsPlayed = strings.Replace(seasonsPlayed, "24", "present", 3)
				}
			}
		}
	}

	if !strings.Contains(seasonsPlayed, "present") {
		seasonsPlayed = strings.ReplaceAll(seasonsPlayed, "-", "-20")
	}

	player := grpc.Player{
		FirstName:     p.FirstName,
		LastName:      p.LastName,
		SeasonsPlayed: seasonsPlayed,
		Seasons:       seasonArray,
	}

	fmt.Println(player)

	playerInfo = append(playerInfo, player)
}

func main() {
	var urls = getPlayerURLs()
	for _, url := range urls {
		scrapePlayerPage(url)
		time.Sleep(5 * time.Second)
	}
	fmt.Println(playerInfo)
	client, err := m.GetMongoClient()
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range playerInfo {
		playerToAdd := bson.D{
			{Key: "firstName", Value: v.FirstName},
			{Key: "lastName", Value: v.LastName},
			{Key: "seasonsPlayed", Value: v.SeasonsPlayed},
			{Key: "seasons", Value: v.Seasons},
		}

		collection := m.GetDatabase(client).Collection("Players")

		_, err = collection.InsertOne(context.TODO(), playerToAdd)
		if err != nil {
			log.Fatal(err)
		}
	}
}
