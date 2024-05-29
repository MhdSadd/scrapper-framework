package extractors

import (
	"log"
	"os"
	"regexp"
	"strings"

	"time"

	"github.com/gocolly/colly"
)

func Extract(url string, htmlSelectors map[string]string) ([]map[string]string) {
		if len(htmlSelectors) == 0 {
			log.Println("❌ No HTML selector (attibutepaths) detected for extractor: at least one selector path is required to proceed")
			os.Exit(1)
		}

			// Use a slice to store data for each attributepaths
			var tempStore []map[string]string

			for key, value := range htmlSelectors {
				log.Println("⏳ mapping through attributepaths")
				c := colly.NewCollector()

				c.SetRequestTimeout(120 * time.Second)

				c.OnRequest(func(r *colly.Request) {
					r.Headers.Set("Accept-Language", "en-US")
					log.Printf("🔗 visiting %s\n", r.URL)
				})

				c.OnError(func(r *colly.Response, err error) {
					log.Println("❌ request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
				})

				re := regexp.MustCompile(`<([^>]+)>`)
				matches := re.FindStringSubmatch(value)
				
				if (len(matches) > 0) {
					value = strings.ReplaceAll(value, matches[0], "")
				}

				c.OnHTML(value, func(h *colly.HTMLElement) {
					result := make(map[string]string)

					if len(matches) < 2 {
						result[key] = strings.TrimSpace(h.Text)
					} else {
						result[key] = h.Attr(matches[1])
					}

					tempStore = append(tempStore, result)
				})

				c.OnScraped(func(r *colly.Response) {
					log.Println("✅ Finished scrapping", r.Request.URL)
				})

				c.Visit(url)
			}

			entry := map[string]string{"source": url}
			tempStore = append(tempStore, entry)
			return tempStore
}
