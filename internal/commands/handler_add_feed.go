package commands

import (
	"Gator/internal/database"
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
)

func Handler_add_feed(s *State, cmd Command, user database.User) error {
	if len(cmd.Args) != 2 {
		return errors.New("You must provide 2 arguments 'Name' and 'URL' of the feed to be added.\n ")
	}
	feed_name := cmd.Args[0]
	feed_url := cmd.Args[1]

	db_feed_params := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      feed_name,
		Url:       feed_url,
		UserID:    user.ID,
	}

	feed, err := s.DB.CreateFeed(context.Background(), db_feed_params)
	if err != nil {
		log.Fatalf("could not create feed: %v", err)
		os.Exit(1)
	}

	db_create_feed_follow_params := database.CreateFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	}

	_, err = s.DB.CreateFeedFollow(context.Background(), db_create_feed_follow_params)
	if err != nil {
		fmt.Printf("Error: couldn't create a feed follow record.\n")
		os.Exit(1)
	}

	fmt.Printf("Feed created successfully: %s (URL: %s | ID: %s)\n", feed.Name, feed.Url, feed.ID)
	return nil
}
