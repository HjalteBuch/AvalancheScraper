package main

import (
	"fmt"
	"strings"
	"time"
    "net/smtp"

	"github.com/gocolly/colly"
)

func main () {
    fmt.Println("AvalancheScraper V0.2")

    now, err := time.Parse("2006.01.02", time.Now().Format("2006.01.02"))
    if err != nil {
        fmt.Println("ERROR:")
        fmt.Println(err)
        panic(err)
    }

    avalancheReport := getAvalancheReport(now)

    if avalancheReport == "" {
        fmt.Println("The avalanche report for today has not been released yet")
        return
    }

    fmt.Println("Retrieved avalanche report for today")

    avalancheReport += "\n Check the lift status on this link: https://www.niseko.ne.jp/en/niseko-lift-status/"

    sendEmail(avalancheReport)
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
            data = ""
        }

    })

    err := c.Visit("http://niseko.nadare.info")
    if err != nil {
        fmt.Println("ERROR:")
        fmt.Println(err)
    }

    return data
}

func sendEmail(message string) {
    from := "hjaltespam@gmail.com"
    password := ""

    to := []string {
        "h.jaltehb123@gmail.com",
    }

    smtpHost := "smtp.gmail.com"
    smtpPort := "587"

    auth := smtp.PlainAuth("", from, password, smtpHost)

    err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, []byte(message))
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println("Email sent")
}
