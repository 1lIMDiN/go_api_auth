package server

import (
	"auth/pkg/api"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Run(port int) error {
	rout := chi.NewRouter()

	api.Init(rout)

	return http.ListenAndServe(fmt.Sprintf(":%d", port), rout)
}