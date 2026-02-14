package commands

import "Gator/internal/config"

func (c *Commands) Register(name string, f func(*config.State, Command) error) {
	c.Cmd_List[name] = f
}
