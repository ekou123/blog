package config

import "fmt"

func HandlerLogin(s *State, cmd Command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("expected a single username argument")
	}

	name := cmd.Args[0]

	err := s.Cfg.SetUser(name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %v", err)
	}

	fmt.Printf("User %s has been set\n", name)

	return nil
}
