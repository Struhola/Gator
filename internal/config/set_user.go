package config

func (c *Config) SetUser(username string) error {
	c.Current_User_Name = username
	return Write(c)
}
