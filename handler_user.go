package main

import (
	"Gator/internal/database"
	"context"
	"fmt"
	"time"
	"github.com/google/uuid"
)

func handlerLogin(s *state, cmd command) error {
	// Validate command arguments
	if len(cmd.Args) != 1 {
		return fmt.Errorf("username is required")
	}
	username := cmd.Args[0]

	user, err := s.db.GetUser(context.Background(), username)
	if err != nil {
		return fmt.Errorf("user does not exist")
	}

	// Attempt to set the user in config
	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("failed to set user: %v", err)
	}

	// Provide success feedback
	fmt.Printf("Successfully logged in as user: %s\n", username)
	return nil
}

func handlerRegister(s *state, cmd command) error {
	// Validate command arguments
	if len(cmd.Args) != 1 {
		return fmt.Errorf("username is required")
	}

	// Check if user already exists
	name := cmd.Args[0]
	_, err := s.db.GetUser(context.Background(), name)
	if err == nil {
		return fmt.Errorf("user already exists")
	}

	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
	})
	if err != nil {
		return fmt.Errorf("failed to create user: %v", err)
	}

	// Set user in config
	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("failed to set user: %v", err)
	}
	return nil

}



func handlerUsers(s *state, cmd command) error {
	// Get users from database
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get users: %v", err)
	}

	// Print users
	for _, user := range users {
		if user.Name == s.cfg.CurrentUserName {
			fmt.Println("* " + user.Name + " (current)")
		} else {
			fmt.Println("* " + user.Name)
		}
	}
	return nil
}
