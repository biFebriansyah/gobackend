package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/biFebriansyah/gobackend/src/libs"
)

func CheckAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		headerToken := r.Header.Get("Authorization")

		if !strings.Contains(headerToken, "Bearer") {
			libs.Respone("invalid header type", 401, true).Send(w)
			return
		}

		token := strings.Replace(headerToken, "Bearer ", "", -1)

		checkTokens, err := libs.CheckToken(token)
		if err != nil {
			libs.Respone(err.Error(), 401, true).Send(w)
			return
		}

		ctx := context.WithValue(r.Context(), "user", checkTokens.User_id)

		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

func AuthWithRole(role ...string) Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-type", "application/json")

			var header string
			var valid bool = false

			if header = r.Header.Get("Authorization"); header == "" {
				libs.Respone("header not provide", 401, false).Send(w)
				return
			}

			if !strings.Contains(header, "Bearer") {
				libs.Respone("invalid header type", 401, false).Send(w)
				return
			}

			token := strings.Replace(header, "Bearer ", "", -1)

			checkTokens, err := libs.CheckToken(token)
			if err != nil {
				libs.Respone(err.Error(), 201, true).Send(w)
				return
			}

			for _, rl := range role {
				if rl == checkTokens.Role {
					valid = true
				}
			}

			if !valid {
				libs.Respone("you not have permission to accsess", 401, false).Send(w)
				return
			}

			// share context to controller
			ctx := context.WithValue(r.Context(), "user", checkTokens.User_id)

			// Serve the next handler
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
