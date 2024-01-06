package main

import (
    "fmt"
	"github.com/gocolly/colly"
	"strconv"
    "strings"
	"time"
    "github.com/gookit/config/v2"
)

type AvalancheData struct {
    Date string
    Text string
}

func main () {
    fmt.Println("AvalancheScraper V0.1")

    days := amountOfDaysSince(28)

    c := colly.NewCollector()

    baseUrl := "http://niseko.nadare.info/?page="

    var data []AvalancheData

    c.OnHTML(".entry", func(e *colly.HTMLElement) {
        text := strings.ToLower(e.ChildText(".entry_body"))
        englishStart := strings.Index(text, "mountain")
        englishEnd := strings.Index(text, "tweet")
        data = append(data, AvalancheData {
            Date: e.ChildText(".entry_date"),
            Text: text[englishStart:englishEnd],
        })
    })

    for i := 0; i < days; i++ {
        url := baseUrl + strconv.Itoa(i)
        err := c.Visit(url)
        if err != nil {
            fmt.Println("ERROR:")
            fmt.Println(err)
        }
    }

    err := config.LoadFiles("internal/config.json")
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(config.Get("APIKey"))
}

func amountOfDaysSince(date int) int {
    currentDate := time.Now()
    fromDate := time.Date(2023, time.December, date, 0, 0, 0, 0, currentDate.Location())
    duration := currentDate.Sub(fromDate)
    return int(duration.Hours() / 24)
}
