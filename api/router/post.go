package router

import (
	"github.com/go-chi/chi"

	"github.com/sithumonline/red-timer/api/handler"
)

func PostRouter() chi.Router {
	r := chi.NewRouter()

	r.Post("/", handler.CreatePost)
	r.Get("/", handler.PostList)
	r.Get("/{id}", handler.GetPost)
	r.Put("/{id}", handler.UpdatePost)
	r.Delete("/{id}", handler.DeletePost)

	return r
}
