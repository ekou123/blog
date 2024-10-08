package config

import (
	"context"
	"fmt"
)

const FeedURL = "https://www.wagslane.dev/index.xml"

func HandlerAgg(s *State, cmd Command) error {

	feed, err := FetchFeed(context.Background(), FeedURL)
	if err != nil {
		return fmt.Errorf("could not fetch feed: %w", err)
	}

	fmt.Printf("%+v\n", feed)

	return nil
}
