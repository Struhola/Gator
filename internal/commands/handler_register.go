package commands

import (
	"Gator/internal/database"
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
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
		if isDuplicateKeyError(err) {
			fmt.Printf("Error: user with name '%s' already exists\n", user_name)
			os.Exit(1)
		}

		log.Fatalf("could not create user: %v", err)
	}

	fmt.Printf("User created successfully: %s (ID: %s)\n", user.Name, user.ID)
	Handler_login(s, cmd)
	return nil
}

func isDuplicateKeyError(err error) bool {
	return err != nil && (strings.Contains(err.Error(), "unique constraint") || strings.Contains(err.Error(), "already exists"))
}
