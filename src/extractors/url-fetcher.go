package extractors

import (
	"log"

	"time"

	"github.com/gocolly/colly"
)

func GetUrls(url string, htmlSelector string) ([]string, error) {
	log.Println("🚀 getting detail urls for extractor...")

	urls := make([]string, 0)

	c := colly.NewCollector()

	c.SetRequestTimeout(120 * time.Second)

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Accept-Language", "en-US")
		log.Printf("🔗 visiting %s\n", r.URL)
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Println("❌ request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	c.OnHTML(htmlSelector, func(h *colly.HTMLElement) {
		if h.Attr("src") != "" {
			urls = append(urls, h.Attr("src"))
		} else if h.Attr("href") != "" {
			urls = append(urls, h.Attr("href"))
		}
	})

	c.OnScraped(func(r *colly.Response) {
		log.Println("✅ finished scrapping", r.Request.URL, "for details URLS")
	})

	c.Visit(url)

	return urls, nil
}
