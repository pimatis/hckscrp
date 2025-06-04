package src

import (
	"fmt"
	"os"

	"github.com/gocolly/colly"
	"github.com/jedib0t/go-pretty/v6/table"
)

type HackerNewsJobs struct {
	Rank     string
	Title    string
	URL      string
	Domain   string
	Time     string
	Comments string
}

func JobsScrape(page int) {
	c := colly.NewCollector()

	var items []HackerNewsJobs

	c.OnHTML("tr.athing", func(e *colly.HTMLElement) {
		item := HackerNewsJobs{}
		item.Rank = e.ChildText("td.title span.rank")
		
		titleElement := e.ChildAttr("td.title span.titleline a", "href")
		item.Title = e.ChildText("td.title span.titleline a")
		item.URL = titleElement
		item.Domain = e.ChildText("td.title span.titleline span.sitebit span.sitestr")

		nextRow := e.DOM.Next()
		if nextRow.Length() > 0 {
			item.Time = nextRow.Find("td.subtext span.age").Text()
			item.Comments = nextRow.Find("td.subtext a").Last().Text()
		}

		items = append(items, item)
	})

	c.OnScraped(func(r *colly.Response) {
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"Rank", "Title", "Domain", "Time", "Comments"})

		for _, item := range items {
			title := item.Title

			if len(title) > 50 {
				title = title[:47] + "..."
			}

			url := "https://" + item.Domain

			t.AppendRow(table.Row{
				item.Rank,
				title,
				url,
				item.Time,
				item.Comments,
			})
		}

		t.Render()
	})

	err := c.Visit("https://news.ycombinator.com/jobs" + fmt.Sprintf("?p=%d", page))
	if err != nil {
		fmt.Println("Error visiting URL:", err)
	}
}
