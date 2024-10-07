package config

import (
	"context"
	"fmt"
)

func HandlerReset(s *State, cmd Command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("no argument expected")
	}

	err := s.Db.DeleteUsers(context.Background())
	if err != nil {
		fmt.Printf("Error deleting users: %v\n", err)
	}

	fmt.Println("Users deleted")
	return nil
}
