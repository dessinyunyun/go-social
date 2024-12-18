package repository

import (
	"context"
	"database/sql"
	"errors"
)

var (
	ErrNotFound = errors.New("Resource not Found")
	ErrConflict = errors.New("resource already exists")
)

type Repository struct {
	Posts interface {
		Create(context.Context, *Post) error
		GetById(context.Context, int64) (*Post, error)
	}

	Users interface {
		Create(context.Context, *User) error
	}
}

func NewRepository(db *sql.DB) Repository {
	return Repository{
		Posts: &PostRepository{db},
		Users: &UserRepository{db},
	}
}
