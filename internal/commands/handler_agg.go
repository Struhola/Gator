package commands

import (
	"Gator/internal/config"
	"Gator/internal/rss"
	"fmt"
	"time"
)

func Handler_agg(s *config.State, cmd Command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <time_between_reqs>", cmd.Name)
	}

	timeBetweenRequests, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("invalid duration: %w", err)
	}

	fmt.Printf("Collecting feeds every %s\n", timeBetweenRequests)

	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		err := rss.Scrape_feeds(s)
		if err != nil {
			fmt.Printf("Error scraping feeds: %v\n", err)
			continue
		}
	}
}
