package api

import (
	"auth/pkg/handlers"

	"github.com/go-chi/chi/v5"
)

func Init(rout *chi.Mux) {
	// Публичные эндпоинты
	rout.Post("/register", handlers.Register)
	rout.Post("/login", handlers.Login)

	// Приватные эндпоинты
	rout.Group(func(r chi.Router) {
		r.Use(handlers.AuthMiddleware)
		r.Get("/documents", handlers.GetDocuments)
	})
}