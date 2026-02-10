package commands

import (
	"Gator/internal/database"
	"context"
	"errors"
	"fmt"
	"os"
)

func Handler_follow(s *State, cmd Command) error {
	if len(cmd.Args) != 1 {
		return errors.New("This function requires a single argument with a feed url.")
	}
	feed_url := cmd.Args[0]
	current_user, err := s.DB.GetUser(context.Background(), s.App_Config.Current_User_Name)
	feed, err := s.DB.GetFeedByUrl(context.Background(), feed_url)
	if err != nil {
		fmt.Printf("Error: couldn't retrieve a feed with URL: %s from database.\n", feed_url)
		os.Exit(1)
	}
	db_create_feed_follow_params := database.CreateFeedFollowParams{
		UserID: current_user.ID,
		FeedID: feed.ID,
	}

	feed_follow, err := s.DB.CreateFeedFollow(context.Background(), db_create_feed_follow_params)
	if err != nil {
		fmt.Printf("Error: couldn't create a feed follow record.\n")
		os.Exit(1)
	}

	fmt.Printf("Feed Followed. Name: %s | User: %s\n", feed_follow.FeedName, feed_follow.UserName)

	return nil
}
