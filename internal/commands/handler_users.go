package commands

import (
	"context"
	"errors"
	"fmt"
	"os"
)

func Handler_users(s *State, cmd Command) error {
	if len(cmd.Args) > 0 {
		return errors.New("This function does not take arguments.")
	}
	users, err := s.DB.GetUsers(context.Background())
	if err != nil {
		fmt.Printf("Error: couldn't retrieve users list.\n")
		os.Exit(1)
	}
	current_user := s.App_Config.Current_User_Name
	for _, u := range users {
		if current_user == u.Name {
			fmt.Printf("* %s (current)\n", u.Name)
		} else {
			fmt.Printf("* %s\n", u.Name)
		}
	}

	return nil
}
