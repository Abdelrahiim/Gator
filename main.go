package main

import (
	"Gator/internal/config"
	"Gator/internal/database"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type state struct {
	cfg *config.Config
	db  *database.Queries
}

func main() {

	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("failed to read config file: %v", err)
		return
	}
	// Import the postgres driver
	db, err := sql.Open("postgres", cfg.DbUrl)
	if err != nil {
		fmt.Printf("failed to open database connection: %v", err)
		return
	}
	dbQueries := database.New(db)
	s := &state{
		cfg: &cfg,
		db:  dbQueries,
	}

	cmds := commands{
		map[string]func(*state, command) error{},
	}

	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)
	cmds.register("agg", handleAgg)
	cmds.register("addfeed", middlewareLoggedIn(handleAddFeed))
	cmds.register("feeds", middlewareLoggedIn(handleFeeds))
	cmds.register("follow", middlewareLoggedIn(handleFeedFollow))
	cmds.register("following", middlewareLoggedIn(handleFollowing))
	cmds.register("unfollow", middlewareLoggedIn(handleFeedUnfollow))

	// Get command line arguments
	args := os.Args

	// Check if any arguments were provided
	if len(args) < 2 {
		fmt.Println("Error: No command provided")
		fmt.Println("Usage: gator <command> [arguments...]")
		os.Exit(1)
	}

	// Create command from arguments
	cmd := command{
		Name: args[1],
		Args: args[2:],
	}

	err = cmds.run(s, cmd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}

}
