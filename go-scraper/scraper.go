package main

import (
	"fmt"
	"strconv"

	"github.com/gocolly/colly"
)

type Country struct {
	Name       string  `json:"name,omitempty"`
	Capital    string  `json:"capital,omitempty"`
	Population int64   `json:"population,omitempty"`
	Area       float64 `json:"area,omitempty"`
}

func main() {
	collector := colly.NewCollector(
		colly.AllowedDomains("www.scrapethissite.com"),
	)

	var countries []Country

	collector.OnHTML("div.country", func(h *colly.HTMLElement) {

		population, err := strconv.ParseInt(h.ChildText("span.country-population"), 10, 64)
		if err != nil {
			fmt.Printf("Error parsing population: %v\n", err)
			population = 0
		}

		area, err := strconv.ParseFloat(h.ChildText("span.country-area"), 64)
		if err != nil {
			fmt.Printf("Error parsing area: %v\n", err)
			area = 0
		}

		country := Country{
			Name:       h.ChildText("h3.country-name"),
			Capital:    h.ChildText("span.country-capital"),
			Population: population,
			Area:       area,
		}

		countries = append(countries, country)
	})

	collector.Visit("https://www.scrapethissite.com/pages/simple/")
	fmt.Println(countries)
}
