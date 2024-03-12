package handlers

import (
	"github.com/eggermarc/go-api/internal/middleware"
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
)

// Capital letter functions are public!
func Handler(r *chi.Mux) {
    r.Use(chimiddle.StripSlashes)

    r.Route("/account", func(router chi.Router) {
        // Authentication
        router.Use(middleware.Authorization)

        router.Get("/balance", GetUserBalance)
    })

}


