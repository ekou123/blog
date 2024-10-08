package config

import (
	"context"
	"database/sql"
	"encoding/xml"
	"fmt"
	"github.com/ekou123/blog/internal/database"
	"github.com/google/uuid"
	"html"
	"io"
	"net/http"
	"time"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func FetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, fmt.Errorf("unable to fetch rss feed: %w", err)
	}

	req.Header.Set("User-Agent", "gator")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("unable to do request: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unable to fetch rss feed with status code: %d", resp.StatusCode)
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("unable to read response body: %w", err)
	}

	var rss RSSFeed
	err = xml.Unmarshal(data, &rss)
	if err != nil {
		return nil, fmt.Errorf("unable to parse rss feed xml: %w", err)
	}

	rss.Channel.Description = html.UnescapeString(rss.Channel.Description)
	rss.Channel.Title = html.UnescapeString(rss.Channel.Title)
	for i := range rss.Channel.Item {
		rss.Channel.Item[i].Title = html.UnescapeString(rss.Channel.Item[i].Title)
		rss.Channel.Item[i].Description = html.UnescapeString(rss.Channel.Item[i].Description)
	}

	return &rss, nil
}

func AddFeed(name string, url string) error {

	data, err := Read()
	if err != nil {
		return fmt.Errorf("unable to read user information: %w", err)
	}

	currentUser := data.User

	db, err := sql.Open("postgres", data.DbURL)
	if err != nil {
		return fmt.Errorf("unable to connect to database: %w", err)
	}
	defer db.Close()

	queries := database.New(db)

	user, err := queries.GetUserByName(context.Background(), currentUser)
	if err != nil {
		return fmt.Errorf("unable to get user: %w", err)
	}

	newFeedStruct := database.CreateFeedParams{
		ID:        uuid.New(),
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Url:       url,
		UserID:    user.ID,
	}

	feed, err := queries.CreateFeed(context.Background(), newFeedStruct)
	if err != nil {
		return fmt.Errorf("unable to create feed: %w", err)
	}

	fmt.Printf("New feed created: %+v\n", feed)

	return nil
}

func HandlerAddFeed(s *State, cmd Command) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("Usage: addfeed <name> <url>")
	}
	return AddFeed(cmd.Args[0], cmd.Args[1])
}

func HandlerListFeeds(s *State, cmd Command) error {
	data, err := Read()
	if err != nil {
		return fmt.Errorf("unable to read user information: %w", err)
	}

	db, err := sql.Open("postgres", data.DbURL)
	if err != nil {
		return fmt.Errorf("unable to connect to database: %w", err)
	}

	queries := database.New(db)

	feedsInfo, err := queries.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("unable to get feeds info: %w", err)
	}

	for i := range feedsInfo {
		userInfo, err := queries.GetUserByID(context.Background(), feedsInfo[i].UserID)
		if err != nil {
			return fmt.Errorf("unable to get user: %w", err)
		}
		fmt.Printf("Name: %s\nURL: %s\nUser Name: %s\n\n", feedsInfo[i].Name, feedsInfo[i].Url, userInfo.Name)
	}

	return nil
}
