package config

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

func HandlerLogin(s *State, cmd Command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("expected a single username argument")
	}

	name := cmd.Args[0]

	_, err := s.Db.GetUser(context.Background(), name)
	if errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("user doesn't exist: %s", name)
	} else if err != nil {
		return fmt.Errorf("error querying user: %v", err)
	}

	err = s.Cfg.SetUser(name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %v", err)
	}

	fmt.Printf("User %s has been set\n", name)

	return nil
}
