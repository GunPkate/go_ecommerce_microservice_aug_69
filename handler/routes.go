package handler

import "github.com/go-chi/chi/v5"

var r *chi.Mux

func RegisterRoutes(handler *handler) *chi.Mux {
	r = chi.NewRouter()

	r.Route("/products", func(r chi.Router) {
		r.Post("/", handler.createProduct)
		r.Get("/", handler.listProducts)

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", handler.getProduct)
			r.Delete("/", handler.deleteProduct)
		})
	})

	return r
}
