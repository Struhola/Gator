package commands

import (
	"Gator/internal/config"
	"fmt"
)

func (c *Commands) Run(s *config.State, cmd Command) error {

	callback, ok := c.Cmd_List[cmd.Name]
	if !ok {
		return fmt.Errorf("error: command '%s' not found", cmd.Name)
	}
	return callback(s, cmd)
}
