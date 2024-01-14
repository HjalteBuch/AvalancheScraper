package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

func main () {
    fmt.Println("AvalancheScraper V0.1")

    now, err := time.Parse("2006.01.02", time.Now().Format("2006.01.02"))
    if err != nil {
        fmt.Println("ERROR:")
        fmt.Println(err)
        panic(err)
    }

    avalancheReport := getAvalancheReport(now)
    fmt.Println(avalancheReport)
}

func getAvalancheReport(now time.Time) string {
    var data string
    c := colly.NewCollector()

    c.OnHTML(".entry", func(e *colly.HTMLElement) {
        dateStr := e.ChildText(".entry_date")
        date, err := time.Parse("2006.01.02", dateStr[0:strings.Index(dateStr, " ")])
        if err != nil {
            fmt.Println("ERROR:")
            fmt.Println(err)
            panic(err)
        }
        if date.Equal(now) {
            text := strings.ToLower(e.ChildText(".entry_body"))
            englishStart := strings.Index(text, "mountain")
            englishEnd := strings.Index(text, "tweet")
            data = text[englishStart:englishEnd]
        } else {
            fmt.Println("Avalanche report has not been released for today yet")
        }

    })

    err := c.Visit("http://niseko.nadare.info")
    if err != nil {
        fmt.Println("ERROR:")
        fmt.Println(err)
    }

    return data
}
