package src

import (
	"fmt"
	"os"
	"strings"

	"github.com/gocolly/colly"
	"github.com/jedib0t/go-pretty/v6/table"
)

type HackerNewsNewComments struct {
	ID      string
	Author  string
	Time    string
	Content string
	Story   string
}

func NewCommentsScrape(page int) {
	c := colly.NewCollector()

	var items []HackerNewsNewComments

	c.OnHTML("tr.athing", func(e *colly.HTMLElement) {
		item := HackerNewsNewComments{}
		item.ID = e.Attr("id")
		item.Author = e.ChildText("span.comhead a.hnuser")
		item.Time = e.ChildText("span.comhead span.age a")
		
		contentElement := e.DOM.Find("div.commtext")
		if contentElement.Length() > 0 {
			content := contentElement.Text()
			content = strings.TrimSpace(content)
			
			if len(content) > 200 {
				content = content[:197] + "..."
			}
			item.Content = content
		}
		
		storyElement := e.DOM.Find("span.onstory a")
		if storyElement.Length() > 0 {
			item.Story = storyElement.Text()
		}

		if item.Author != "" && item.Content != "" {
			items = append(items, item)
		}
	})

	c.OnScraped(func(r *colly.Response) {
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"ID", "Author", "Time", "Content", "Story"})

		for _, item := range items {
			content := item.Content
			if len(content) > 80 {
				content = content[:77] + "..."
			}
			
			story := item.Story
			if len(story) > 50 {
				story = story[:47] + "..."
			}

			t.AppendRow(table.Row{
				item.ID,
				item.Author,
				item.Time,
				content,
				story,
			})
		}

		t.Render()
	})

	err := c.Visit("https://news.ycombinator.com/newcomments" + fmt.Sprintf("?p=%d", page))
	if err != nil {
		fmt.Println("Error visiting URL:", err)
	}
}