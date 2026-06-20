package http_server

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	models "yat-blog-notes/internal/infra/http_server/models"
	response "yat-blog-notes/internal/infra/http_server/response"
)

func isAuthenticated(f http.HandlerFunc) http.HandlerFunc {
	// TODO: переделать на jwt (сейчас для дебага норм)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("User-Id")
		if header == "" {
			response.WriteErrorResponse(
				w,
				"user not authenticated",
				http.StatusUnauthorized,
			)
			return
		}
		userId, err := strconv.Atoi(header)
		if err != nil {
			response.WriteErrorResponse(
				w,
				fmt.Sprintf("error parsing `User-Id` header: %v", err),
				http.StatusUnauthorized,
			)
			return
		}
		ctx := context.WithValue(
			r.Context(),
			models.UserCtxKey{},
			models.User{Id: userId},
		)
		f.ServeHTTP(w, r.WithContext(ctx))
	})
}
