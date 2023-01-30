package handler

import (
	"github.com/gmhafiz/go8/internal/domain/account/usecase"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

func RegisterHTTPEndPoints(router *chi.Mux, validate *validator.Validate, useCase usecase.Account) *Handler {
	h := NewHandler(useCase, validate)

	router.Route("/api/v1/account", func(router chi.Router) {
		router.Post("/", h.Create)
	})

	return h
}
