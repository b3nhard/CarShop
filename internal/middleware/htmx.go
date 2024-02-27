package middleware

import (
	"context"
	"net/http"
)

type HTMX string

// Define your context key
const (
	HTMXRequest HTMX = "htmx"
)

func HTMXMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if the request has the HX-Request header
		if h := r.Header.Get("HX-Request"); h != "" {
			// If it does, set the context key to true
			ctx := context.WithValue(r.Context(), HTMXRequest, true)
			// Pass the updated context to the next handler
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}
		// If not, set the context key to false
		ctx := context.WithValue(r.Context(), HTMXRequest, false)
		// Pass the updated context to the next handler
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
