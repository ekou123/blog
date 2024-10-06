package config

func (c *Commands) Register(name string, f func(*State, Command) error) {
	if c.Handler == nil {
		c.Handler = map[string]func(*State, Command) error{}
	}

	c.Handler[name] = f
}
