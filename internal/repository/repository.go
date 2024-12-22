package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

var (
	ErrNotFound          = errors.New("resource not Found")
	ErrConflict          = errors.New("resource already exists")
	QueryTimeoutDuration = time.Second * 5
)

type Repository struct {
	Posts interface {
		Create(context.Context, *Post) error
		GetPost(context.Context, int64) (*Post, error)
		Update(context.Context, *Post) error
		Delete(context.Context, int64) error
	}

	Users interface {
		Create(context.Context, *User) error
		GetUser(context.Context, int64) (*User, error)
	}

	Comments interface {
		Create(context.Context, *Comment) error
		GetByPostId(context.Context, int64) ([]Comment, error)
	}
}

func NewRepository(db *sql.DB) Repository {
	return Repository{
		Posts:    &PostRepository{db},
		Users:    &UserRepository{db},
		Comments: &CommentsRepository{db},
	}
}
