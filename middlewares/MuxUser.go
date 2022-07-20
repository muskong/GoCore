package middlewares

import (
	"net/http"

	"github.com/muskong/GoService/pkg/work"
)

func UserMiddleware(fn http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path

		notAuth := map[string]bool{
			"/user/login": true,
		}
		if !notAuth[path] {
			work.WorkNew(w, r)
			if work.Context.CheckAuth() {
				return
			}
		}
		fn.ServeHTTP(w, r)
	})
}
