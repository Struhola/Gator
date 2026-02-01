package Commands

import (
	"errors"
	"fmt"
)

func Handler_Login(s *State, cmd Command) error {
	if len(cmd.Args) != 1 {
		return errors.New("You must provide a single argument of a user name.")
	}
	User_Name := cmd.Args[0]
	s.App_Config.SetUser(User_Name)
	fmt.Printf("User: %s has been set.\n", User_Name)

	return nil
}
