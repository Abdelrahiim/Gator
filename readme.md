# gator

gator is a command-line RSS feed aggregator and reader. It allows users to register, log in, and manage their own collection of RSS feeds. Users can add new feeds, follow or unfollow feeds, and browse the latest posts from the feeds they follow. Gator periodically fetches and updates feed content, storing posts in a PostgreSQL database for easy retrieval and browsing.

## Prerequisites
- **Go** (version 1.24.5 or higher)
- **PostgreSQL** (running and accessible)

## Installation

1. **Clone the repository:**
   ```sh
   git clone https://github.com/your-username/gator.git
   cd gator
   ```

2. **Install the gator CLI:**
   ```sh
   go install ./...
   ```
   This will build and install the `gator` binary to your `$GOPATH/bin`.

## Configuration

Before running Gator, you need to set up a configuration file in your home directory:

1. Create a `.gatorconfig.json` file in your home directory (e.g., `/home/youruser/.gatorconfig.json`):
   ```json
   {
     "db_url": "postgres://username:password@localhost:5432/gatordb?sslmode=disable",
     "current_user_name": ""
   }
   ```
   - Replace `username`, `password`, and `gatordb` with your actual PostgreSQL credentials and database name.

2. Make sure your PostgreSQL server is running and the database exists.

## Usage

Run the CLI using:
```sh
gator <command> [arguments...]
```

### Example Commands
- **Register a new user:**
  ```sh
  gator register <username>
  ```
- **Login as a user:**
  ```sh
  gator login <username>
  ```
- **Add a new feed:**
  ```sh
  gator addfeed <feed-name> <feed-url>
  ```
- **List all feeds:**
  ```sh
  gator feeds
  ```
- **Follow a feed:**
  ```sh
  gator follow <feed-url>
  ```
- **Browse posts:**
  ```sh
  gator browse [limit]
  ```

## Pushing to GitHub

1. Initialize a git repository if you haven't already:
   ```sh
   git init
   git add .
   git commit -m "Initial commit"
   ```
2. Create a new repository on GitHub and follow the instructions to add the remote:
   ```sh
   git remote add origin https://github.com/your-username/gator.git
   git push -u origin main
   ```

## Submit Your Repo

Submit your repository link in the following format:
```
https://github.com/your-username/gator
```
