package rss

import (
	"Gator/internal/config"
	"Gator/internal/database"
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/google/uuid"
)

func Scrape_feeds(s *config.State) error {
	feed, err := s.DB.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't get next feed: %w", err)
	}

	rss_feed, err := Fetch_feed(context.Background(), feed.Url)
	if err != nil {
		return fmt.Errorf("couldn't fetch feed: %s | URL: %s | Err: %w", feed.Url, feed.Name, err)
	}

	err = s.DB.MarkFeedFetched(context.Background(), database.MarkFeedFetchedParams{
		ID:        feed.ID,
		UpdatedAt: time.Now().UTC(),
	})

	if err != nil {
		return fmt.Errorf("couldn't mark feed as fetched: %w", err)
	}

	fmt.Printf("Found %d items in feed %s\n", len(rss_feed.Channel.Item), feed.Name)
	for _, item := range rss_feed.Channel.Item {
		publishedAt := sql.NullTime{}
		if t, err := Parse_date(item.PubDate); err == nil {
			publishedAt = sql.NullTime{
				Time:  t,
				Valid: true,
			}
		}

		_, err = s.DB.CreatePost(context.Background(), database.CreatePostParams{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			Title:     item.Title,
			Url:       item.Link,
			Description: sql.NullString{
				String: item.Description,
				Valid:  item.Description != "",
			},
			PublishedAt: publishedAt,
			FeedID:      feed.ID,
		})

		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				continue
			}
			log.Printf("Couldn't create post: %v", err)
		}
	}
	log.Printf("Feed %s collected, %v posts found", feed.Name, len(rss_feed.Channel.Item))

	return nil
}
