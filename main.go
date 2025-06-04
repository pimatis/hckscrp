package main

import (
	"fmt"
	src "hckscrp/src/home"
	"hckscrp/src/user"
)

func main() {
	fmt.Println("Welcome to HCKSCRP - Hacker News Scraper")
	
	for {
		fmt.Println("Available commands:")
		fmt.Println("1. Front Page")
		fmt.Println("2. News")
		fmt.Println("3. Ask")
		fmt.Println("4. Show")
		fmt.Println("5. Jobs")
		fmt.Println("6. New Comments")
		fmt.Println("7. Fetch Ask Comments")
		fmt.Println("8. User Submissions")
		fmt.Println("9. User Threads")
		fmt.Println("10. User Favorites")
		fmt.Println("11. User Info")
		fmt.Println("12. Exit")
		fmt.Println("-----------------------------------------")
		var command int
		fmt.Print("Enter command (1-12): ")
		fmt.Scanln(&command)

		switch command {
		case 1:
			var page int
			fmt.Print("Enter page number (default is 1): ")
			fmt.Scanln(&page)
			if page <= 0 {
				page = 1
			}
			fmt.Println("Fetching Hacker News Front Page...")
			src.FrontScrape(page)
		case 2:
			var page int
			fmt.Print("Enter page number (default is 1): ")
			fmt.Scanln(&page)
			if page <= 0 {
				page = 1
			}
			fmt.Println("Fetching Hacker News...")
			src.NewsScrape(page)
		case 3:
			var page int
			fmt.Print("Enter page number (default is 1): ")
			fmt.Scanln(&page)
			if page <= 0 {
				page = 1
			}
			fmt.Println("Fetching Ask HN...")
			src.AskScrape(page)
		case 4:
			var page int
			fmt.Print("Enter page number (default is 1): ")
			fmt.Scanln(&page)
			if page <= 0 {
				page = 1
			}
			fmt.Println("Fetching Show HN...")
			src.ShowScrape(page)
		case 5:
			var page int
			fmt.Print("Enter page number (default is 1): ")
			fmt.Scanln(&page)
			if page <= 0 {
				page = 1
			}
			fmt.Println("Fetching Hacker News Jobs...")
			src.JobsScrape(page)
		case 6:
			var page int
			fmt.Print("Enter page number (default is 1): ")
			fmt.Scanln(&page)
			if page <= 0 {
				page = 1
			}
			fmt.Println("Fetching New Comments on Hacker News...")
			src.NewCommentsScrape(page)
		case 7:
			var itemID string
			fmt.Print("Enter Item ID for Ask Comments: ")
			fmt.Scanln(&itemID)
			fmt.Printf("Fetching comments for item ID: %s\n", itemID)
			src.FetchAskComments(itemID)
		case 8:
			var username string
			var page int
			fmt.Print("Enter username: ")
			fmt.Scanln(&username)
			fmt.Print("Enter page number (default is 1): ")
			fmt.Scanln(&page)
			if page <= 0 {
				page = 1
			}
			fmt.Printf("Fetching submissions for user: %s\n", username)
			user.SubmittedScrape(username, page)
		case 9:
			var username string
			var page int
			fmt.Print("Enter username: ")
			fmt.Scanln(&username)
			fmt.Print("Enter page number (default is 1): ")
			fmt.Scanln(&page)
			if page <= 0 {
				page = 1
			}
			fmt.Printf("Fetching threads for user: %s\n", username)
			user.ThreadsScrape(username, page)
		case 10:
			var username string
			var page int
			fmt.Print("Enter username: ")
			fmt.Scanln(&username)
			fmt.Print("Enter page number (default is 1): ")
			fmt.Scanln(&page)
			if page <= 0 {
				page = 1
			}
			fmt.Printf("Fetching favorites for user: %s\n", username)
			user.FavoritesScrape(username, page)
		case 11:
			var username string
			fmt.Print("Enter username: ")
			fmt.Scanln(&username)
			fmt.Printf("Fetching user info for: %s\n", username)
			user.UserInfoScrape(username)
		case 12:
			fmt.Println("Exiting hckscrp...")
			return
		default:
			fmt.Println("Invalid command. Please enter a number between 1-12.")
		}
		fmt.Println()
	}
}
