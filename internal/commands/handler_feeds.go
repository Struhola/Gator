package commands

import (
	"context"
	"errors"
	"fmt"
	"os"
)

func Handler_feeds(s *State, cmd Command) error {
	if len(cmd.Args) > 0 {
		return errors.New("This function does not take arguments.")
	}
	feeds, err := s.DB.GetFeeds(context.Background())
	if err != nil {
		fmt.Printf("Error: couldn't retrieve feeds list.\n")
		os.Exit(1)
	}
	//current_user := s.App_Config.Current_User_Name
	fmt.Printf("All available feeds:\n")
	for _, f := range feeds {
		fmt.Printf("* Name: %s | URL: %s | Owner: %s\n", f.Name, f.Url, f.UserName)
	}

	return nil
}
