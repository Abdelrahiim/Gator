package main

import (
	"Gator/internal/database"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func handleFeedFollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("url is required")
	}

	url := cmd.Args[0]

	feed, err := s.db.GetFeed(context.Background(), url)
	if err != nil {
		return fmt.Errorf("failed to get feed: %v", err)
	}

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		UserID:    user.ID,
		FeedID:    feed.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return fmt.Errorf("failed to create feed follow: %v", err)
	}

	fmt.Printf("Feed follow created successfully:\n"+
		"ID: %v\n"+
		"User ID: %v\n"+
		"Feed ID: %v\n"+
		"Created At: %v\n"+
		"Updated At: %v\n",
		feedFollow.ID,
		feedFollow.UserID,
		feedFollow.FeedID,
		feedFollow.CreatedAt.Format(time.RFC3339),
		feedFollow.UpdatedAt.Format(time.RFC3339))
	return nil
}

func handleFeedUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("url is required")
	}
	url := cmd.Args[0]
	feed, err := s.db.GetFeed(context.Background(), url)
	if err != nil {
		return fmt.Errorf("failed to get feed: %v", err)
	}

	err = s.db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return fmt.Errorf("failed to delete feed follow: %v", err)
	}
	return nil
}

func handleFollowing(s *state, cmd command, user database.User) error {
	feedFollowing, err := s.db.GetFeedFollowsForUser(context.Background(), user.Name)
	if err != nil {
		return fmt.Errorf("failed to get feed follow: %v", err)
	}
	for _, feedFollow := range feedFollowing {
		fmt.Printf("\t * Feed Name: %-20s\n\t   Feed URL: %-30s\n\t   Created by: %s\n\t   Created At: %s\n\t   Updated At: %s\n\t   ID: %s\n\n",
			feedFollow.FeedName,
			feedFollow.FeedUrl,
			feedFollow.UserName,
			feedFollow.CreatedAt.Format(time.RFC3339),
			feedFollow.UpdatedAt.Format(time.RFC3339),
			feedFollow.ID)
	}
	return nil

}
