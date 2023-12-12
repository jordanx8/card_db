package main

import (
	"fmt"
	"log"
	"strings"

	// importing Colly
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

type Listing struct {
	Title   string
	Price   string
	Details map[string]string
}

var listings []Listing

func hasStrikethrough(s string) bool {
	return strings.Contains(s, "STRIKETHROUGH")
}

func main() {
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
				Price:   price,
				Details: deets,
			}
			listings = append(listings, l)
		}
	})

	err := c.Visit("https://www.ebay.com/sch/i.html?_from=R40&_trksid=p2334524.m570.l1313&_nkw=2022-23+Mosaic+Thunder+Road+%2312+Stephen+Curry+Warriors+Insert&_sacat=0&LH_TitleDesc=0&_fsrp=1&rt=nc&Player%252FAthlete=Stephen%2520Curry&LH_Complete=1&LH_Sold=1&Sport=Basketball&_odkw=2022-23+Mosaic+Thunder+Road+%2312+Stephen+Curry+Warriors+Insert+base&_osacat=0&_dcat=261328&Team=Golden%2520State%2520Warriors&Features=Insert")
	if err != nil {
		log.Printf("failed to visit url: %v\n", err)
		return
	}

	fmt.Println(listings)
	fmt.Println(len(listings))
}
