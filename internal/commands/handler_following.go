package commands

import (
	"context"
	"errors"
	"fmt"
	"os"
)

func Handler_following(s *State, cmd Command) error {
	if len(cmd.Args) > 0 {
		return errors.New("This function does not take arguments.")
	}

	feeds, err := s.DB.GetFeedFollowsForUser(context.Background(), s.App_Config.Current_User_Name)
	if err != nil {
		fmt.Printf("Error: couldn't retrieve feeds list.\n")
		os.Exit(1)
	}
	current_user := s.App_Config.Current_User_Name
	fmt.Printf("All feeds followed by %s:\n", current_user)
	for _, f := range feeds {
		fmt.Printf("* Name: %s\n", f.FeedName)
	}

	return nil
}
