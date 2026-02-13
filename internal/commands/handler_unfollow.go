package commands

import (
	"Gator/internal/database"
	"context"
	"errors"
	"fmt"
	"os"
)

func Handler_unfollow(s *State, cmd Command, user database.User) error {
	if len(cmd.Args) != 1 {
		return errors.New("This function requires a single argument with a feed url.")
	}
	feed_url := cmd.Args[0]
	feed, err := s.DB.GetFeedByUrl(context.Background(), feed_url)
	if err != nil {
		fmt.Printf("Error: couldn't retrieve a feed with URL: %s from database.\n", feed_url)
		os.Exit(1)
	}
	db_create_feed_follow_params := database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	}

	err = s.DB.DeleteFeedFollow(context.Background(), db_create_feed_follow_params)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		fmt.Printf("Error while clearing the table.")
		os.Exit(1)
	}

	fmt.Printf("Feed Unfollowed. Name: %s | User: %s\n", feed_url, user.Name)

	return nil
}
