package commands

import (
	"Gator/internal/database"
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
)

func Handler_register(s *State, cmd Command) error {
	if len(cmd.Args) != 1 {
		return errors.New("You must provide a single argument of a user name.")
	}
	user_name := cmd.Args[0]
	db_user_params := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      user_name,
	}
	user, err := s.DB.CreateUser(context.Background(), db_user_params)
	if err != nil {
		log.Fatalf("could not create user: %v", err)
		os.Exit(1)
	}

	fmt.Printf("User created successfully: %s (ID: %s)\n", user.Name, user.ID)
	err = s.App_Config.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}
	return nil
}
