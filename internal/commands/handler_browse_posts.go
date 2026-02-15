package commands

import (
	"Gator/internal/config"
	"Gator/internal/database"
	"context"
	"fmt"
	"strconv"
)

func Handler_browse(s *config.State, cmd Command, user database.User) error {
	limit := 2
	if len(cmd.Args) > 0 {
		if l, err := strconv.Atoi(cmd.Args[0]); err == nil {
			limit = l
		} else {
			return fmt.Errorf("invalid limit: %w", err)
		}
	}

	posts, err := s.DB.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		return fmt.Errorf("could check for posts: %w", err)
	}

	fmt.Printf("Found %d posts for user %s:\n", len(posts), user.Name)
	for _, post := range posts {
		publishedAt := "Unknown Date"
		if post.PublishedAt.Valid {
			publishedAt = post.PublishedAt.Time.Format("Feb 05, 2026")
		}

		fmt.Printf("--- %s ---\n", post.Title)
		fmt.Printf("Source: %s\n", post.Url)
		fmt.Printf("Published: %s\n", publishedAt)
		if post.Description.Valid && post.Description.String != "" {
			fmt.Printf("Description: %s\n", post.Description.String)
		}
		fmt.Println()
	}

	return nil
}
