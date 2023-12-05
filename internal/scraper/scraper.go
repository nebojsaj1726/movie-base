package scraper

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

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
	var totalPagesMutex sync.Mutex

	listCollector := colly.NewCollector()
	movieCollector := colly.NewCollector()

	listCollector.OnHTML(".movie-item-style-2", func(e *colly.HTMLElement) {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()

			title := e.ChildText(".mv-item-infor h6 a")
			rate := e.ChildText(".rate span")
			year := e.ChildText(".year")

			link := e.ChildAttr("a[href]", "href")

			movie := Movie{
				Title: title,
				Rate: rate,
				Year: year,
			}

			err := movieCollector.Visit(e.Request.AbsoluteURL(link))
			if err != nil {
				log.Printf("Error visiting page %s: %v\n", link, err)
				return
			}

			movies = append(movies, movie)
		}(&wg)		
	})

	movieCollector.OnHTML("#overview", func(e *colly.HTMLElement) {
		genres := extractGenres(e)
		duration := e.ChildText(".movie-description__duration span")
		description := e.ChildText(".description")

		mu.Lock()
		defer mu.Unlock()

		lastMovieIndex := len(movies) - 1
		if lastMovieIndex >= 0 {
			movies[lastMovieIndex].Genres = genres
			movies[lastMovieIndex].Duration = duration
			movies[lastMovieIndex].Description = description
		}
	})

	listCollector.OnHTML(".pagination__right", func(e *colly.HTMLElement) {
		totalPagesMutex.Lock()
    	defer totalPagesMutex.Unlock()

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
			log.Printf("Total Pages: %d\n", totalPages)
		}
	})

	listCollector.OnError(func(r *colly.Response, err error) {
		log.Printf("Request URL: %s failed with response: %v\n", r.Request.URL, r)
	})

	movieCollector.OnError(func(r *colly.Response, err error) {
		log.Printf("Request URL: %s failed with response: %v\n", r.Request.URL, r)
	})

	log.Printf("Total Pages: %d\n", totalPages)
	startUrl := "https://www.lookmovie2.to/page/%d?&r=5"
	for i := 1; i <= 1; i++ {
		err := listCollector.Visit(fmt.Sprintf(startUrl, i))
		if err != nil {
			log.Printf("Error visiting page %d: %v\n", i, err)
		}

		time.Sleep(1 * time.Second)
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