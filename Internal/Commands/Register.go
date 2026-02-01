package Commands

func (c *Commands) Register(name string, f func(*State, Command) error) {
	c.Cmd_List[name] = f
}
