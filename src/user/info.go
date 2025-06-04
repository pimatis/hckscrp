package user

import (
	"fmt"
	"os"

	"github.com/gocolly/colly"
	"github.com/jedib0t/go-pretty/v6/table"
)

type HackerNewsUserInfo struct {
	Username string
	Created  string
	Karma    string
	About    string
}

func UserInfoScrape(username string) {
	c := colly.NewCollector()

	var userInfo HackerNewsUserInfo

	c.OnHTML("table tr.athing", func(e *colly.HTMLElement) {
		userInfo.Username = e.ChildText("td a.hnuser")
	})

	c.OnHTML("table tr", func(e *colly.HTMLElement) {
		label := e.ChildText("td:first-child")
		value := e.ChildText("td:nth-child(2)")

		switch label {
		case "created:":
			userInfo.Created = value
		case "karma:":
			userInfo.Karma = value
		case "about:":
			userInfo.About = value
		}
	})

	c.OnScraped(func(r *colly.Response) {
		if userInfo.Username == "" {
			fmt.Printf("User '%s' not found or profile is private.\n", username)
			return
		}

		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"Field", "Value"})

		t.AppendRow(table.Row{"Username", userInfo.Username})
		t.AppendRow(table.Row{"Created", userInfo.Created})
		t.AppendRow(table.Row{"Karma", userInfo.Karma})
		
		if userInfo.About != "" {
			t.AppendRow(table.Row{"About", userInfo.About})
		}

		t.Render()
	})

	url := fmt.Sprintf("https://news.ycombinator.com/user?id=%s", username)
	err := c.Visit(url)
	if err != nil {
		fmt.Println("Error visiting URL:", err)
	}
}
