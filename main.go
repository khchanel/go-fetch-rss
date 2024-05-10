package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/mmcdole/gofeed"
)

func main() {

	urls := []string{
		"https://www.smh.com.au/rss/feed.xml",
		"https://www.smh.com.au/rss/politics/federal.xml",
		"https://www.smh.com.au/rss/national/nsw.xml",
		"https://www.smh.com.au/rss/world.xml",
		"https://www.smh.com.au/rss/national.xml",
		"https://www.smh.com.au/rss/business.xml",
		"https://www.smh.com.au/rss/culture.xml",
		"https://www.smh.com.au/rss/technology.xml",
		"https://www.smh.com.au/rss/environment.xml",
		"https://www.smh.com.au/rss/lifestyle.xml",
		"https://www.smh.com.au/rss/property.xml",
		"https://www.smh.com.au/rss/goodfood.xml",
		"https://www.smh.com.au/rss/traveller.xml",
		"https://www.smh.com.au/rss/sport.xml",
		"https://www.smh.com.au/rss/sport/nrl.xml",
		"https://www.smh.com.au/rss/sport/rugby-union.xml",
		"https://www.smh.com.au/rss/sport/afl.xml",
		"https://rss.nytimes.com/services/xml/rss/nyt/HomePage.xml",
	}

	var wg sync.WaitGroup
	ch := make(chan *gofeed.Feed)

	// start goroutines to fetch RSS feeds concurrently
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			fetchRSS(url, ch)
		}(url)
	}

	// close channel after all goroutines finish
	go func() {
		wg.Wait()
		close(ch)
	}()

	// print feeds
	for feed := range ch {
		printFeed(*feed)
	}

}

func fetchRSS(url string, ch chan<- *gofeed.Feed) {

	fp := gofeed.NewParser()

	feed, err := fp.ParseURL(url)
	if err != nil {
		log.Printf("Error fetching or parsing RSS feed (%s): %v", url, err)
		return
	}

	ch <- feed
}

func printFeed(feed gofeed.Feed) {

	fmt.Println("Feed Title:", feed.Title)
	fmt.Println("Feed Description:", feed.Description)
	fmt.Println("Feed Link:", feed.Link)

	fmt.Println("\nItems:")
	for _, item := range feed.Items {
		fmt.Println("Title:", item.Title)
		fmt.Println("Description:", item.Description)
		fmt.Println("Link:", item.Link)
		fmt.Println("Published Date:", item.Published)
		fmt.Println()
	}
}
