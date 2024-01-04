package main

import (
    "fmt"
	"github.com/gocolly/colly"
	"strconv"
	"time"
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
        data = append(data, AvalancheData {
            Date: e.ChildText(".entry_date"),
            Text: e.ChildText(".entry_body"),
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

}

func amountOfDaysSince(date int) int {
    currentDate := time.Now()
    fromDate := time.Date(2023, time.December, date, 0, 0, 0, 0, currentDate.Location())
    duration := currentDate.Sub(fromDate)
    return int(duration.Hours() / 24)
}
