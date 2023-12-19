package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

type Player struct {
	FirstName string
	LastName  string
}

type Listing struct {
	Title   string
	Price   float64
	Details map[string]string
}

type CardSearch struct {
	Year     string
	Player   Player
	Parallel string
	Set      string
	Team     string
	Insert   string
	CardNo   string
}

func hasStrikethrough(s string) bool {
	return strings.Contains(s, "STRIKETHROUGH")
}

func convertPriceToFloat(price string) float64 {
	price = strings.Replace(price, "US $", "", -1)
	if s, err := strconv.ParseFloat(price, 64); err == nil {
		return s
	}
	return 0
}

func scrapeCompletedListings(ebayUrl string) []Listing {
	var listings []Listing

	c := colly.NewCollector(
		colly.AllowedDomains("ebay.com", "www.ebay.com"),
	)

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong: ", err)
	})

	c.OnHTML("li.s-item.s-item__pl-on-bottom", func(e *colly.HTMLElement) {
		listingurl := e.ChildAttr("a", "href")
		c.Visit(listingurl)
	})

	c.OnHTML("div.main-container", func(e *colly.HTMLElement) {
		if !hasStrikethrough(e.ChildAttr("div.x-price-primary > span.ux-textspans", "class")) {
			var deets = make(map[string]string)

			title := e.ChildText("h1.x-item-title__mainTitle")
			price := e.ChildText("div.x-price-primary")

			e.DOM.Find("div.ux-layout-section-evo__col").Each(func(_ int, s *goquery.Selection) {
				label := s.Find("div.ux-labels-values__labels").Text()
				value := s.Find("div.ux-labels-values__values").Text()
				if label != "" || value != "" {
					deets[label] = value
				}
			})
			l := Listing{
				Title:   title,
				Price:   convertPriceToFloat(price),
				Details: deets,
			}
			listings = append(listings, l)
		}
	})

	err := c.Visit(ebayUrl)
	if err != nil {
		log.Printf("failed to visit url: %v\n", err)
		return nil
	}

	return listings
}

func createEbayURL(cs CardSearch) string {
	s := cs.Year + " " + cs.Set + " "
	if cs.Insert != "" {
		s = s + cs.Insert + " Insert "
	}
	s = s + "#" + cs.CardNo + " " + cs.Player.FirstName + " " + cs.Player.LastName + " " + cs.Team
	urlStart := "https://www.ebay.com/sch/i.html?_from=R40&_nkw="
	urlEnd := "&_sacat=0&rt=nc&LH_Sold=1&LH_Complete=1"
	s = strings.Replace(s, " ", "+", -1)
	s = strings.Replace(s, "#", "%23", -1)
	return urlStart + s + urlEnd
}

func validationsNeeded(cs CardSearch) []func(Listing, CardSearch) bool {
	var validations []func(Listing, CardSearch) bool
	validations = append(validations, playerNameAndYearValidation, setValidation)
	if cs.Insert != "" {
		validations = append(validations, insertValidation)
	}
	if cs.Parallel == "Base" {
		validations = append(validations, baseValidation)
	} else {
		validations = append(validations, parallelValidation)
	}
	return validations
}

func filterListings(listings []Listing, cs CardSearch) []Listing {
	var filteredListings []Listing
	var validations = validationsNeeded(cs)
ListingLoop:
	for _, listing := range listings {
		for _, validation := range validations {
			if !validation(listing, cs) {
				continue ListingLoop
			}
		}
		filteredListings = append(filteredListings, listing)

	}
	return filteredListings
}

func createSearch(year string, playerName string, set string, team string, parallel string, insert string, cardno string) CardSearch {
	name := strings.Split(playerName, " ")
	return CardSearch{
		Year: year,
		Player: Player{
			FirstName: name[0],
			LastName:  name[1],
		},
		Parallel: parallel,
		Set:      set,
		Team:     team,
		Insert:   insert,
		CardNo:   cardno,
	}
}

func main() {
	cardParameters := createSearch(
		"2022-23",
		"Matt Ryan",
		"Panini Instant",
		"Los Angeles Lakers",
		"",
		"",
		"24",
	)
	searchURL := createEbayURL(cardParameters)
	fmt.Println(searchURL)
	listings := scrapeCompletedListings(searchURL)
	filtered := filterListings(listings, cardParameters)
	fmt.Println(cardParameters)
	fmt.Println("--------------------------------")
	if len(filtered) == 0 {
		fmt.Println("No recent purchases.")
	} else {
		sum := 0.0
		high := filtered[0].Price
		low := filtered[0].Price
		for _, v := range filtered {
			if v.Price < low {
				low = v.Price
			}
			if v.Price > high {
				high = v.Price
			}
			sum += v.Price
		}
		fmt.Println("# of Listings: ", len(filtered))
		fmt.Println("Total: ", sum)
		fmt.Println("Average: ", sum/float64(len(filtered)))
		fmt.Println("High: ", high)
		fmt.Println("Low: ", low)
	}

	// write to csv for testing
	csvFile, err := os.Create("listings2.csv")
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	defer csvFile.Close()

	w := csv.NewWriter(csvFile)
	defer w.Flush()

	for _, listing := range filtered {
		row := []string{listing.Title, strconv.FormatFloat(listing.Price, 'f', -1, 64)}
		for k, v := range listing.Details {
			row = append(row, k+": "+v)
		}
		if err := w.Write(row); err != nil {
			log.Fatalln("error writing record to file", err)
		}
	}
}
