package http_server

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	response "yat-blog-notes/internal/infra/http_server/response"
)

func (nc *NotesController) GetAuthorNotes(w http.ResponseWriter, req *http.Request) {
	q := req.URL.Query()

	page := 1
	limit := 20

	if v, err := strconv.Atoi(q.Get("page")); err == nil {
		page = v
	}
	if v, err := strconv.Atoi(q.Get("limit")); err == nil {
		limit = v
	}

	authorId, err := strconv.Atoi(q.Get("author_id"))
	if err != nil || authorId <= 0 {
		response.WriteErrorResponse(w, "invalid author_id", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(
		req.Context(),
		time.Second*time.Duration(nc.config.Http.Timeout),
	)
	defer cancel()

	notes, err := nc.notesSrv.GetAuthorNotes(ctx, authorId, page, limit)
	if err != nil {
		response.WriteErrorResponse(
			w,
			fmt.Sprintf("error while getting notes: %v", err),
			http.StatusBadRequest,
		)
		return
	}
	response.WriteResponse(w, notes, http.StatusOK)
}
