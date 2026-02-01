package Commands

import (
	"Gator/Internal/Config"
)

type State struct {
	App_Config *Config.Config
}

type Command struct {
	Name string
	Args []string
}

type Commands struct {
	Cmd_List map[string]func(*State, Command) error
}
