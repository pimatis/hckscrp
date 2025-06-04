package src

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

type HackerNewsComment struct {
	ID      string
	Author  string
	Time    string
	Content string
	Indent  string
}

func FetchAskComments(itemID string) {
	c := colly.NewCollector()

	var comments []HackerNewsComment

	c.OnHTML("tr.athing.comtr", func(e *colly.HTMLElement) {
		comment := HackerNewsComment{}
		comment.ID = e.Attr("id")

		indentImg := e.ChildAttr("td.ind img", "width")
		if indentImg != "" {
			if width, err := strconv.Atoi(indentImg); err == nil {
				comment.Indent = fmt.Sprintf("%d", width/40)
			}
		} else {
			comment.Indent = "0"
		}

		comment.Author = e.ChildText("a.hnuser")
		comment.Time = e.ChildText("span.age a")
		contentElement := e.DOM.Find("div.commtext")
		if contentElement.Length() > 0 {
			contentElement.Find("form").Remove()
			contentElement.Find("textarea").Remove()
			contentElement.Find("input").Remove()

			content := contentElement.Text()
			content = strings.TrimSpace(content)

			if len(content) > 200 {
				content = content[:197] + "..."
			}
			comment.Content = content
		}

		if comment.Author != "" && comment.Content != "" {
			comments = append(comments, comment)
		}
	})

	c.OnScraped(func(r *colly.Response) {
		for _, comment := range comments {
			fmt.Printf("ID: %s\n", comment.ID)
			fmt.Printf("Level: %s\n", comment.Indent)
			fmt.Printf("Author: %s\n", comment.Author)
			fmt.Printf("Time: %s\n", comment.Time)
			fmt.Printf("Content: %s\n", comment.Content)
			fmt.Println("---")
		}
	})

	err := c.Visit("https://news.ycombinator.com/item?id=" + itemID)
	if err != nil {
		fmt.Println("Error visiting URL:", err)
	}
}