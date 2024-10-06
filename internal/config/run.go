package config

import (
	"errors"
	"fmt"
)

func (c *Commands) Run(s *State, cmd Command) error {
	command, ok := c.Handler[cmd.Name]
	if !ok {
		return errors.New("Command not found")
	}

	err := command(s, cmd)
	if err != nil {
		return fmt.Errorf("Error executing command: %v", err)
	}

	return nil
}
