package middlewarejwt

import (
	"net/http"
)

func RequirePermission(moduleID uint, perm string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			modules, ok := r.Context().Value(ContextModulesKey).(map[uint][]string)
			if !ok {
				http.Error(w, "forbidden", http.StatusForbidden)
				return
			}

			perms := modules[moduleID]
			for _, p := range perms {
				if p == perm {
					next.ServeHTTP(w, r)
					return
				}
			}

			http.Error(w, "forbidden", http.StatusForbidden)
		})
	}
}
