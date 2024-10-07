package config

import (
	"context"
	"fmt"
)

func HandlerUsers(s *State, cmd Command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("not expecting any arguments")
	}

	allUsers, err := s.Db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("Error getting users: %w\n", err)
	}

	for _, user := range allUsers {
		line := fmt.Sprintf("* %s", user)
		if user == s.Cfg.User {
			line += " (current)"
		}
		fmt.Println(line)

	}

	return nil
}
