package main

import (
    "fmt"
	"github.com/gocolly/colly"
    "strings"
	"time"
)

func main () {
    fmt.Println("AvalancheScraper V0.1")

    var data string

    now, err := time.Parse("2006.01.02", time.Now().Format("2006.01.02"))
    if err != nil {
        fmt.Println("ERROR:")
        fmt.Println(err)
        panic(err)
    }

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

    baseUrl := "http://niseko.nadare.info"
    err = c.Visit(baseUrl)
    if err != nil {
        fmt.Println("ERROR:")
        fmt.Println(err)
    }

    fmt.Println(data)
}
