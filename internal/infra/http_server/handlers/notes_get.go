package http_server

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	response "yat-blog-notes/internal/infra/http_server/response"
)

func (nc *NotesController) GetNote(w http.ResponseWriter, req *http.Request) {
	noteId, err := strconv.Atoi(req.PathValue("id"))
	if err != nil {
		// TODO: replace with TryProcessSlug(w, req)
		response.WriteErrorResponse(w, "not found", http.StatusNotFound)
		return
	}

	ctx, cancel := context.WithTimeout(
		req.Context(),
		time.Duration(nc.config.Http.Timeout)*time.Second,
	)
	defer cancel()

	note, err := nc.notesSrv.GetNote(ctx, noteId)
	if err != nil {
		response.WriteErrorResponse(
			w,
			fmt.Sprintf("error while fetching note: %v", err),
			http.StatusNotFound,
		)
		return
	}
	response.WriteResponse(w, ToResponseGetNote(note), http.StatusOK)
}
