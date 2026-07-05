package http_server

import (
	"context"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"time"

	response "yat-blog-notes/internal/infra/http_server/response"
)

func (nc *NotesController) GetNote(w http.ResponseWriter, req *http.Request) {
	noteId, err := strconv.Atoi(req.PathValue("id"))
	if err != nil {
		nc.tryProcessSlug(w, req)
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

var slugPattern = regexp.MustCompile(`^[a-z0-9]+(?:-[a-z0-9]+)*$`)

func (nc *NotesController) tryProcessSlug(w http.ResponseWriter, req *http.Request) {
	slug := req.PathValue("id")

	if slug == "" || !slugPattern.MatchString(slug) {
		response.WriteErrorResponse(w, "not found", http.StatusNotFound)
		return
	}

	ctx, cancel := context.WithTimeout(
		req.Context(),
		time.Duration(nc.config.Http.Timeout)*time.Second,
	)
	defer cancel()

	note, err := nc.notesSrv.GetNoteBySlug(ctx, slug)
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
