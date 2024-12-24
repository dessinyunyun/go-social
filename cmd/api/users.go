package main

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dessinyunyun/socialgo/internal/repository"
	"github.com/go-chi/chi/v5"
)

type userKey string

const userCtx userKey = "user"

func (app *application) getUserHandler(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.ParseInt(chi.URLParam(r, "userId"), 10, 64)
	fmt.Println("tesss", userId)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	ctx := r.Context()

	user, err := app.repository.Users.GetUser(ctx, userId)

	if err != nil {
		switch err {
		case repository.ErrNotFound:
			app.badRequestResponse(w, r, err)
			return
		default:
			app.internalServerError(w, r, err)
			return
		}
	}

	if err := app.jsonResponse(w, http.StatusOK, user); err != nil {
		app.internalServerError(w, r, err)
	}
}

type Followuser struct {
	UserId int64 `json:"user_id"`
}

func (app *application) followUserHandler(w http.ResponseWriter, r *http.Request) {
	followerUser := getUserFromContext(r)

	var payload Followuser
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	ctx := r.Context()

	if err := app.repository.Followers.Follow(ctx, payload.UserId, followerUser.ID); err != nil {
		switch err {
		case repository.ErrConflict:
			app.conflictResponse(w, r, err)
			return
		default:
			app.internalServerError(w, r, err)
			return
		}
	}

	if err := app.jsonResponse(w, http.StatusNoContent, nil); err != nil {
		app.internalServerError(w, r, err)
	}
}

func (app *application) unfollowUserHandler(w http.ResponseWriter, r *http.Request) {
	unfollowuser := getUserFromContext(r)

	var payload Followuser
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	ctx := r.Context()

	if err := app.repository.Followers.Unfollow(ctx, payload.UserId, unfollowuser.ID); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusNoContent, nil); err != nil {
		app.internalServerError(w, r, err)
	}
}

func (app *application) userContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userId, err := strconv.ParseInt(chi.URLParam(r, "userId"), 10, 64)

		if err != nil {
			app.badRequestResponse(w, r, err)
			return
		}

		ctx := r.Context()

		user, err := app.repository.Users.GetUser(ctx, userId)
		if err != nil {
			switch err {
			case repository.ErrNotFound:
				app.notFoundResponse(w, r, err)
				return
			default:
				app.internalServerError(w, r, err)
				return
			}
		}

		ctx = context.WithValue(ctx, userCtx, user)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getUserFromContext(r *http.Request) *repository.User {
	user, _ := r.Context().Value(userCtx).(*repository.User)
	return user
}
