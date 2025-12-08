package middleware

import "net/http"

func RequirePermission(p string) func(http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			claims := r.Context().Value(UserCtxKey).(*UserClaims)

			ok := false
			for _, perm := range claims.Permissions {
				if perm == p {
					ok = true
					break
				}
			}

			if !ok {
				http.Error(w, "forbidden", http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
