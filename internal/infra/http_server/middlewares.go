package http_server

import (
	"net/http"

	response "yat-blog-notes/internal/infra/http_server/response"
)

func isAuthenticated(f http.HandlerFunc) http.HandlerFunc {
	// TODO: переделать на jwt (сейчас для дебага норм)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userId := r.Header.Get("User-Id")
		if userId == "" {
			response.WriteErrorResponse(
				w,
				"user not authenticated",
				http.StatusUnauthorized,
			)
			return
		}

		f.ServeHTTP(w, r)
	})
}
