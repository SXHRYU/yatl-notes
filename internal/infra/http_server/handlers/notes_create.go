package http_server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	models "yat-blog-notes/internal/infra/http_server/models"
	response "yat-blog-notes/internal/infra/http_server/response"
)

func (ac *NotesController) CreateNote(rw http.ResponseWriter, req *http.Request) {
	// todo: implement using structs and `validator`
	const size = 10 << 10

	user := req.Context().Value(models.UserCtxKey{})
	body := http.MaxBytesReader(rw, req.Body, size)
	defer body.Close()

	var r struct{ Text string }

	if err := json.NewDecoder(body).Decode(&r); err != nil {
		if _, ok := errors.AsType[*http.MaxBytesError](err); ok {
			response.WriteErrorResponse(
				rw,
				fmt.Sprintf("`text` size is bigger than allowed (%d KiB)", size>>10),
				http.StatusRequestEntityTooLarge,
			)
			return
		}
		response.WriteErrorResponse(
			rw,
			fmt.Sprintf("failed to decode body: %v", err),
			http.StatusBadRequest,
		)
		return
	}

	ctx, cancel := context.WithTimeout(
		req.Context(),
		time.Duration(ac.config.Http.Timeout)*time.Second,
	)
	defer cancel()

	noteId, err := ac.notesSrv.CreateNote(ctx, user.(models.User).Id, r.Text)
	if err != nil {
		response.WriteErrorResponse(
			rw,
			fmt.Sprintf("error while creating note: %v", err),
			http.StatusBadRequest,
		)
		return
	}

	response.WriteResponse(rw, ToResponseCreateNote(noteId), http.StatusCreated)
}
