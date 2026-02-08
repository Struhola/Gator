package commands

import (
	"context"
	"errors"
	"fmt"
	"os"
)

func Handler_reset(s *State, cmd Command) error {
	if len(cmd.Args) > 0 {
		return errors.New("This function takes no arguments.")
	}
	err := s.DB.ClearUsers(context.Background())
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		fmt.Printf("Error while clearing the table.")
		os.Exit(1)
	}
	return nil
}
