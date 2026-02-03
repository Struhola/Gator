package commands

import (
	"context"
	"errors"
	"fmt"
	"os"
)

func Handler_login(s *State, cmd Command) error {
	if len(cmd.Args) != 1 {
		return errors.New("You must provide a single argument of a user name.")
	}
	user_name := cmd.Args[0]
	_, err := s.DB.GetUser(context.Background(), user_name)
	if err != nil {
		fmt.Printf("Error: user with name '%s' not found in the database.\n", user_name)
		os.Exit(1)
	}
	s.App_Config.SetUser(user_name)
	fmt.Printf("User: %s has been set.\n", user_name)

	return nil
}
