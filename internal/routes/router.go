package routes

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	get_counters_handler "otus_sn_counters/internal/handlers/get-counters-handler"
	update_counters_handler "otus_sn_counters/internal/handlers/update-counters-handler"
)

func RegisterRouter() *chi.Mux {
	r := chi.NewRouter()
	r.With(
		RequestIDMiddleware,
		cors.Handler(cors.Options{
			// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
			AllowedOrigins: []string{"https://*", "http://*", "*"},
			// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: false,
			MaxAge:           300, // Maximum value not ignored by any of major browsers
		}),
	).Route("/api", func(r chi.Router) {
		r.Route("/counters", func(r chi.Router) {
			r.Post("/{id}", update_counters_handler.UpdateCountersHandler)
			r.Get("/{id}", get_counters_handler.GetCountersHandler)
		})
	})
	return r
}
