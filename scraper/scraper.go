package scraper

import (
	"easyfitanalysis/analyser"
	"encoding/csv"
	"github.com/gocolly/colly"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type UsageEntry struct {
	timestamp int
	usage     int
}

// ScrapePage Scrapes the page for the usage data
func ScrapePage(channel chan<- interface{}) {
	url := "https://easyfitness.club/studio/easyfitness-luebeck/"
	var ScrapedUsage int

	c := colly.NewCollector()

	c.OnRequest(func(request *colly.Request) {
		log.Println("Scraping page: ", url)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("No data scrapable", err)
		ScrapedUsage = -1
	})

	c.OnHTML(".meterbubble", func(element *colly.HTMLElement) {
		usageString := element.Text

		ScrapedUsage, _ = strconv.Atoi(strings.TrimSuffix(usageString, "%"))

	})

	c.OnScraped(func(r *colly.Response) {
		log.Println("Finished scraping...")
	})

	c.Visit(url)

	if ScrapedUsage > 100 || ScrapedUsage <= 0 {
		ScrapedUsage = -1
	}

	entry := UsageEntry{
		timestamp: int(time.Now().Unix()),
		usage:     ScrapedUsage,
	}

	err := writeUsage(entry)
	if err != nil {
		log.Panicf("There was an error writing the scraped data: %s", err)
	}

	analysis := analyser.ReturnAnalysis()

	channel <- analysis

	return
}

// Appends the usage to the csv-File
func writeUsage(entry UsageEntry) (err error) {
	log.Println("Attempting csv write...")

	f, err := os.OpenFile("./util_log.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()

	if err != nil {
		return err
	}

	w := csv.NewWriter(f)

	writableEntry := []string{strconv.Itoa(entry.timestamp), strconv.Itoa(entry.usage)}

	w.Write(writableEntry)
	w.Flush()

	log.Println("Wrote to csv...")
	return
}
