package router

import (
	"github.com/go-chi/chi"

	"github.com/sithumonline/red-timer/api/handler"
)

func TimeRouter() chi.Router {
	r := chi.NewRouter()

	r.Post("/", handler.ResetTime)
	r.Get("/", handler.GetTime)
	r.Get("/ui", handler.Ui)

	return r
}
