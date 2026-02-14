package rss

import (
	"Gator/internal/config"
	"Gator/internal/database"
	"context"
	"fmt"
	"time"
)

func ScrapeFeeds(s *config.State) error {
	feed, err := s.DB.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't get next feed: %w", err)
	}

	err = s.DB.MarkFeedFetched(context.Background(), database.MarkFeedFetchedParams{
		ID:        feed.ID,
		UpdatedAt: time.Now().UTC(),
	})
	if err != nil {
		return fmt.Errorf("couldn't mark feed as fetched: %w", err)
	}

	rss_feed, err := FetchFeed(context.Background(), feed.Url)
	if err != nil {
		return fmt.Errorf("couldn't fetch feed: %s | URL: %s | Err: %w", feed.Url, feed.Name, err)
	}

	fmt.Printf("Found %d items in feed %s\n", len(rss_feed.Channel.Item), feed.Name)
	for _, item := range rss_feed.Channel.Item {
		fmt.Printf("- %s\n", item.Title)
	}

	return nil
}
