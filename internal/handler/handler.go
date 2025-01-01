package handler

import (
	"github.com/Ibrahimkazi18/First-Go/internal/middleware"
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
)

func Handler(r *chi.Mux) {
	r.Use(chimiddle.StripSlashes)

	r.Route("/account", func(router chi.Router) {

		//Middleware for /acount route
		router.Use(middleware.Authorization)

		// router.Get("/coins", GetCoinBalance)
	})
}
