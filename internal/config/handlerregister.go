package config

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/ekou123/blog/internal/database"
	"github.com/google/uuid"
	"time"
)

func HandlerRegister(s *State, cmd Command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("expected a single username argument")
	}

	username := cmd.Args[0]

	_, err := s.Db.GetUser(context.Background(), username)
	if !errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("user %s already exists", username)
	}

	newUserStruct := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      username,
	}

	user, err := s.Db.CreateUser(context.Background(), newUserStruct)
	if err != nil {
		return fmt.Errorf("cannot register user: %w", err)
	}

	err = s.Cfg.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("cannot set user: %w", err)
	}

	fmt.Printf("User registered: %s\n", user.Name)

	return nil

}
