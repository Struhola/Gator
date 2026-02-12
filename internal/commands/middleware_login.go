package commands

import (
	"Gator/internal/database"
	"context"
	"fmt"
)

func Middleware_logged_in(handler func(s *State, cmd Command, user database.User) error) func(*State, Command) error {
	return func(s *State, cmd Command) error {
		userName := s.App_Config.Current_User_Name
		if userName == "" {
			return fmt.Errorf("this command requires you to be logged in")
		}

		user, err := s.DB.GetUser(context.Background(), userName)
		if err != nil {
			return fmt.Errorf("could not find user %s: %w", userName, err)
		}

		return handler(s, cmd, user)
	}
}
