package commands

import (
	"Gator/internal/config"
	"context"
	"errors"
	"fmt"
	"os"
)

func Handler_reset(s *config.State, cmd Command) error {
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
