package user

import (
	"fmt"
	"os"
	"strings"

	"github.com/gocolly/colly"
	"github.com/jedib0t/go-pretty/v6/table"
)

type HackerNewsFavorites struct {
	Rank     string
	Title    string
	URL      string
	Domain   string
	Score    string
	Author   string
	Time     string
	Comments string
}

func FavoritesScrape(username string, page int) {
	c := colly.NewCollector()

	var items []HackerNewsFavorites
	var nextURL string

	c.OnHTML("tr.athing", func(e *colly.HTMLElement) {
		item := HackerNewsFavorites{}
		item.Rank = e.ChildText("td.title span.rank")
		
		titleElement := e.ChildAttr("td.title span.titleline a", "href")
		item.Title = e.ChildText("td.title span.titleline a")
		item.URL = titleElement
		item.Domain = e.ChildText("td.title span.titleline span.sitebit span.sitestr")

		nextRow := e.DOM.Next()
		if nextRow.Length() > 0 {
			item.Score = nextRow.Find("td.subtext span.score").Text()
			item.Author = nextRow.Find("td.subtext a.hnuser").Text()
			item.Time = nextRow.Find("td.subtext span.age a").Text()
			commentsText := nextRow.Find("td.subtext a").Last().Text()
			if strings.Contains(commentsText, "comment") {
				item.Comments = commentsText
			} else {
				item.Comments = "discuss"
			}
		}

		items = append(items, item)
	})

	c.OnHTML("a.morelink", func(e *colly.HTMLElement) {
		nextURL = e.Attr("href")
	})

	c.OnScraped(func(r *colly.Response) {
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"Rank", "Title", "Domain", "Score", "Author", "Time", "Comments"})

		for _, item := range items {
			title := item.Title
			if len(title) > 50 {
				title = title[:47] + "..."
			}

			url := "https://" + item.Domain
			if item.Domain == "" && item.URL != "" {
				if strings.HasPrefix(item.URL, "http") {
					url = item.URL
				} else {
					url = "https://news.ycombinator.com/" + item.URL
				}
			}

			t.AppendRow(table.Row{
				item.Rank,
				title,
				url,
				item.Score,
				item.Author,
				item.Time,
				item.Comments,
			})
		}

		t.Render()

		if nextURL != "" && page > 1 {
			fmt.Printf("\nFetching next page...\n")
			err := c.Visit("https://news.ycombinator.com/" + nextURL)
			if err != nil {
				fmt.Printf("Error visiting next page: %v\n", err)
			}
		}
	})

	var url string
	if page == 1 {
		url = fmt.Sprintf("https://news.ycombinator.com/favorites?id=%s", username)
	} else {
		url = fmt.Sprintf("https://news.ycombinator.com/favorites?id=%s&p=%d", username, page)
	}

	err := c.Visit(url)
	if err != nil {
		fmt.Println("Error visiting URL:", err)
	}
}
