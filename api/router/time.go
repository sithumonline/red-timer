package router

import (
	"github.com/go-chi/chi"

	"github.com/sithumonline/red-timer/api/handler"
)

func TimeRouter() chi.Router {
	r := chi.NewRouter()

	//r.Post("/", handler.CreatePuso)
	r.Get("/", handler.GetTime)
	//r.Get("/{id}", handler.GetPuso)
	//r.Put("/{id}", handler.UpdatePuso)
	//r.Delete("/{id}", handler.DeletePuso)

	return r
}
