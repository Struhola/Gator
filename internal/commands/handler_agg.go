package commands

import (
	"Gator/internal/rss"
	"context"
	"fmt"
)

func Handler_agg(s *State, cmd Command) error {
	feedURL := "https://www.wagslane.dev/index.xml"

	feed, err := rss.FetchFeed(context.Background(), feedURL)
	if err != nil {
		return fmt.Errorf("couldn't fetch feed: %w", err)
	}

	fmt.Printf("%+v\n", feed)
	return nil
}
