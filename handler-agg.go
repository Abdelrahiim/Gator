package main

import (
	"Gator/internal/database"
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

func handleAgg(s *state, cmd command) error {
	if len(cmd.Args) < 1 || len(cmd.Args) > 2 {
		return fmt.Errorf("usage: %v <time_between_reqs>", cmd.Name)
	}

	timeBetweenRequests, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("invalid duration: %w", err)
	}

	ticker := time.NewTicker(timeBetweenRequests)

	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

func scrapeFeeds(s *state) {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		fmt.Printf("Error fetching feed: %v\n", err)
		return
	}
	feed, err = s.db.MarkAsFetched(context.Background(), feed.ID)

	if err != nil {
		fmt.Printf("Error marking feed as fetched: %v\n", err)
		return
	}

	realFeed, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		fmt.Printf("Error fetching feed: %v\n", err)
		return
	}

	for _, item := range realFeed.Channel.Item {
		publishedAt := sql.NullTime{}
		if t, parseErr := time.Parse(time.RFC1123Z, item.PubDate); parseErr == nil {
			publishedAt = sql.NullTime{
				Time:  t,
				Valid: true,
			}
		}
		_, err = s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       item.Title,
			Url:         item.Link,
			Description: sql.NullString{String: item.Description, Valid: true},
			PublishedAt: publishedAt,
			FeedID:      feed.ID,
		})
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				continue
			}
			fmt.Printf("Error creating post: %v\n", err)
			return
		}

	}
}
