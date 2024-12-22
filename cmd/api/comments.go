package main

import (
	"net/http"

	"github.com/dessinyunyun/socialgo/internal/repository"
)

type createCommentPayload struct {
	PostId  int64  `json:"post_id"`
	UserId  int64  `json:"user_id"`
	Content string `json:"content"`
}

func (app *application) createCommentHandler(w http.ResponseWriter, r *http.Request) {
	var payload createCommentPayload
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	comment := &repository.Comment{
		PostId:  int64(payload.PostId),
		UserId:  int64(payload.UserId),
		Content: payload.Content,
	}

	ctx := r.Context()

	if err := app.repository.Comments.Create(ctx, comment); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusCreated, comment); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}
