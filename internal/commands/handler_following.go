package commands

import (
	"Gator/internal/database"
	"context"
	"errors"
	"fmt"
	"os"
)

func Handler_following(s *State, cmd Command, user database.User) error {
	if len(cmd.Args) > 0 {
		return errors.New("This function does not take arguments.")
	}

	feeds, err := s.DB.GetFeedFollowsForUser(context.Background(), user.Name)
	if err != nil {
		fmt.Printf("Error: couldn't retrieve feeds list.\n")
		os.Exit(1)
	}

	fmt.Printf("All feeds followed by %s:\n", user.Name)
	for _, f := range feeds {
		fmt.Printf("* Name: %s\n", f.FeedName)
	}

	return nil
}
