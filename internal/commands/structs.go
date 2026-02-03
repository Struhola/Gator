package commands

import (
	"Gator/internal/config"
	"Gator/internal/database"
)

type State struct {
	DB         *database.Queries
	App_Config *config.Config
}

type Command struct {
	Name string
	Args []string
}

type Commands struct {
	Cmd_List map[string]func(*State, Command) error
}
