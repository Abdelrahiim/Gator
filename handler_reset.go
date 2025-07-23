package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	// Reset user in config
	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		return fmt.Errorf("failed to reset user: %v", err)
	}
	return nil
}
