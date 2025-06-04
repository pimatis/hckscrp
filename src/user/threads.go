package user

import (
	"fmt"
	"os"
	"strings"

	"github.com/gocolly/colly"
	"github.com/jedib0t/go-pretty/v6/table"
)

type HackerNewsThreads struct {
	ID       string
	Author   string
	Time     string
	Content  string
	Story    string
	StoryURL string
}

func ThreadsScrape(username string, page int) {
	c := colly.NewCollector()

	var items []HackerNewsThreads
	var nextURL string

	c.OnHTML("tr.athing", func(e *colly.HTMLElement) {
		item := HackerNewsThreads{}
		item.ID = e.Attr("id")
		item.Author = e.ChildText("span.comhead a.hnuser")
		item.Time = e.ChildText("span.comhead span.age a")
		
		contentElement := e.DOM.Find("div.commtext")
		if contentElement.Length() > 0 {
			content := contentElement.Text()
			content = strings.TrimSpace(content)
			
			if len(content) > 300 {
				content = content[:297] + "..."
			}
			item.Content = content
		}
		
		storyElement := e.DOM.Find("span.onstory a")
		if storyElement.Length() > 0 {
			item.Story = storyElement.Text()
			item.StoryURL = storyElement.AttrOr("href", "")
			if item.StoryURL != "" && !strings.HasPrefix(item.StoryURL, "http") {
				item.StoryURL = "https://news.ycombinator.com/" + item.StoryURL
			}
		}

		if item.Author != "" && item.Content != "" {
			items = append(items, item)
		}
	})

	c.OnHTML("a.morelink", func(e *colly.HTMLElement) {
		nextURL = e.Attr("href")
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
		url = fmt.Sprintf("https://news.ycombinator.com/threads?id=%s", username)
	} else {
		url = fmt.Sprintf("https://news.ycombinator.com/threads?id=%s&p=%d", username, page)
	}

	err := c.Visit(url)
	if err != nil {
		fmt.Println("Error visiting URL:", err)
	}
}
