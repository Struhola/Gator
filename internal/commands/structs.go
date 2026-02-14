package commands

import "Gator/internal/config"

type Command struct {
	Name string
	Args []string
}

type Commands struct {
	Cmd_List map[string]func(*config.State, Command) error
}
