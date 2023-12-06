package scraper

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"sync"

	"github.com/gocolly/colly/v2"
)

type Movie struct {
	Title       string
	Rate        string
	Year        string
	Description string
	Genres      []string
	Duration    string
}

func ScrapeMovies() ([]Movie, error) {
	var movies []Movie
	var wg sync.WaitGroup
	var mu sync.Mutex
	var totalPages int

	listCollector := colly.NewCollector()
	pageCollector := colly.NewCollector()


	wg.Add(1)
	go func() {
		defer wg.Done()

		listCollector.OnHTML(".pagination__right", func(e *colly.HTMLElement) {
			totalPagesStr := e.Text
			re := regexp.MustCompile(`Page 1 of (\d+)`)
			matches := re.FindStringSubmatch(totalPagesStr)
		
			var err error
		
			if len(matches) >= 2 {
				totalPages, err = strconv.Atoi(matches[1])
				if err != nil {
					log.Printf("Error converting total pages: %v\n", err)
					return
				}
			}
		})

		listCollector.OnHTML(".movie-item-style-2", func(e *colly.HTMLElement) {
			title := e.ChildText(".mv-item-infor h6 a")
			rate := e.ChildText(".rate span")
			year := e.ChildText(".year")
			link := e.ChildAttr("a[href]", "href")

			movie := Movie{
				Title: title,
				Rate: rate,
				Year: year,
			}

			overviewCollector := colly.NewCollector()

			overviewCollector.OnHTML("#overview", func(e *colly.HTMLElement) {
				genres := extractGenres(e)
				duration := e.ChildText(".movie-description__duration span")
				description := e.ChildText(".description")
		
				mu.Lock()
				defer mu.Unlock()
		
				movie.Genres = genres
				movie.Duration = duration
				movie.Description = description

				movies = append(movies, movie)
			})

			err := overviewCollector.Visit(e.Request.AbsoluteURL(link))
			if err != nil {
				log.Printf("Error visiting page %s: %v\n", link, err)
				return
			}	
		})

		listCollector.OnError(func(r *colly.Response, err error) {
			log.Printf("Request URL: %s failed with response: %v\n", r.Request.URL, r)
		})

		err := listCollector.Visit("https://www.lookmovie2.to/page/1?&r=5")
		if err != nil {
			log.Printf("Error visiting page 1: %v\n", err)
		}
	}()

	wg.Wait()
	log.Printf("Total Pages: %d\n", totalPages)
	for i := 1; i <= 1; i++ {
		wg.Add(1)
		go func(page int) {
			defer wg.Done()

			err := pageCollector.Visit(fmt.Sprintf("https://www.lookmovie2.to/page/%d?&r=5", page))
			if err != nil {
				log.Printf("Error visiting page %d: %v\n", page, err)
			}
		}(i)
	}

	wg.Wait()

	return movies, nil
}

func extractGenres(e *colly.HTMLElement) []string {
	var genres []string
	e.ForEach(".genres span", func(_ int, el *colly.HTMLElement) {
		genre := strings.TrimSpace(el.Text)
		genres = append(genres, genre)
	})
	return genres
}