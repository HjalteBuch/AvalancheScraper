package main

import (
    "fmt"
	"github.com/gocolly/colly"
	"strconv"
    "strings"
	"time"
    "github.com/gookit/config/v2"
    // "net/http"
)

type Gptrequest struct {
    Prompt string
    Data string
}

func main () {
    fmt.Println("AvalancheScraper V0.1")

    err := config.LoadFiles("config.json")
    if err != nil {
        fmt.Println(err)
    }
    days := amountOfDaysSince(config.Int("days"))
    baseUrl := config.String("baseUrl")

    c := colly.NewCollector()

    var data string

    c.OnHTML(".entry", func(e *colly.HTMLElement) {
        text := strings.ToLower(e.ChildText(".entry_body"))
        englishStart := strings.Index(text, "mountain")
        englishEnd := strings.Index(text, "tweet")
        data = data + e.ChildText(".entry_date") + text[englishStart:englishEnd] + "\n\n"
    })

    for i := 0; i < days; i++ {
        url := baseUrl + strconv.Itoa(i)
        err := c.Visit(url)
        if err != nil {
            fmt.Println("ERROR:")
            fmt.Println(err)
        }
    }
    fmt.Println(data)

    // Contact chatGPT
    // req, err := http.NewRequest("POST", config.String("APIUrl"), data)
    // req.Header.Add("Content-Type", "application/json")
    // req.Header.Add("Authorization", config.String("APIKey"))
}

func amountOfDaysSince(date int) int {
    currentDate := time.Now()
    fromDate := time.Date(2023, time.December, date, 0, 0, 0, 0, currentDate.Location())
    duration := currentDate.Sub(fromDate)
    return int(duration.Hours() / 24)
}
