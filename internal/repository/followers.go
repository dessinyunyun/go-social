package repository

import (
	"context"
	"database/sql"

	"github.com/lib/pq"
)

type Follower struct {
	UserId     int64  `json:"user_id"`
	FollowerId int64  `json:"follower_id"`
	CreatedAt  string ` json:"created_at"`
}

type FollowerRepository struct {
	db *sql.DB
}

func (s *FollowerRepository) Follow(ctx context.Context, userId, followerId int64) error {
	query := `INSERT INTO followers (user_id, follower_id) VALUES ($1, $2)`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	_, err := s.db.ExecContext(ctx, query, userId, followerId)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			return ErrConflict
		}
	}

	return nil
}

func (s *FollowerRepository) Unfollow(ctx context.Context, userId, followerId int64) error {
	query := `DELETE FROM followers
			  WHERE user_id = $1 AND follower_id = $2`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	_, err := s.db.ExecContext(ctx, query, userId, followerId)
	return err
}
