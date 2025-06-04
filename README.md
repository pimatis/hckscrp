# HCKSCRP - Hacker News Scraper

A comprehensive command-line tool for scraping various sections of Hacker News, built with Go and the Colly web scraping framework.

## Features

### Home Page Scrapers
- **Front Page** - Scrape the main Hacker News front page
- **News** - Scrape the latest news stories
- **Ask HN** - Scrape Ask Hacker News posts
- **Show HN** - Scrape Show Hacker News posts
- **Jobs** - Scrape job postings
- **New Comments** - Scrape the latest comments across the site
- **Fetch Ask Comments** - Get detailed comments for specific Ask HN posts

### User-Specific Scrapers
- **User Submissions** - View all posts submitted by a specific user
- **User Threads** - View all comments made by a specific user
- **User Favorites** - View posts favorited by a specific user
- **User Info** - Get detailed profile information for any user

## Installation

1. Clone the repository:
```bash
git clone https://github.com/pimatis/hckscrp.git
cd hckscrp
```

2. Install dependencies:
```bash
go mod tidy
```

3. Build the application:
```bash
go build -o hckscrp main.go
```

## Usage

Run the application:
```bash
./hckscrp
```

or if you prefer to run it with `go run`:
```bash
go run main.go
```

You'll be presented with an interactive menu:

```
Welcome to HCKSCRP - Hacker News Scraper
Available commands:
1. Front Page
2. News
3. Ask
4. Show
5. Jobs
6. New Comments
7. Fetch Ask Comments
8. User Submissions
9. User Threads
10. User Favorites
11. User Info
12. Exit
-----------------------------------------
Enter command (1-12):
```

### Examples

#### Scraping Home Pages
- Select options 1-6 and enter a page number when prompted
- Page numbers start from 1 (default)

#### Fetching Comments
- Select option 7 and enter a specific item ID
- Example: `44178902` for a specific Ask HN post

#### User-Specific Operations
- Select options 8-11 and enter a username when prompted
- For submissions, threads, and favorites, you can also specify a page number
- Example username: `queaxtra`

## Output Format

All scraped data is displayed in formatted tables with relevant columns:

### Story Tables Include:
- Rank
- Title (truncated to 50 characters)
- Domain/URL
- Score (when available)
- Author
- Time posted
- Comment count

### Comment Tables Include:
- Comment ID
- Author
- Time posted
- Content (truncated for readability)
- Story context

### User Info Table Includes:
- Username
- Account creation date
- Karma score
- About section (if available)

## Dependencies

- [Colly](https://github.com/gocolly/colly) - Web scraping framework
- [go-pretty](https://github.com/jedib0t/go-pretty) - Table formatting

## Features

- **Pagination Support**: All relevant scrapers support multiple pages
- **Automatic URL Handling**: Properly formats both external and internal Hacker News links
- **Error Handling**: Graceful handling of network errors and missing data
- **Clean Output**: Formatted tables with appropriate column widths and text truncation
- **Interactive Menu**: Easy-to-use command-line interface

## Rate Limiting

Please be respectful when using this scraper:
- Don't make too many rapid requests
- Consider adding delays between requests for heavy usage
- Follow Hacker News' robots.txt and terms of service

## Contributing

Feel free to submit issues and enhancement requests!

## License

This project is open source and available under the [MIT License](LICENSE).
