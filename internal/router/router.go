package router

import (
	"context"
	"net/http"

	"github.com/b3nhard/car-rental/web/components"
	"github.com/go-chi/chi/v5"
)

func SetupRoutes(r *chi.Mux) {

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		components.Layout("Home", components.Index()).Render(context.Background(), w)
	})
	r.NotFound(http.NotFound)
}
