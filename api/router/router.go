package router

import "github.com/go-chi/chi"

func Router() chi.Router {
	r := chi.NewRouter()

	r.Mount("/time", TimeRouter())

	return r
}
