package main

import "github.com/go-chi/chi"

func RegisterRoutes(router *chi.Mux) {
	router.Get("/api/blogs", GetBlogs)
	// router.Get("/api/blogs/{id}", getBlog)
	// router.Post("/api/blogs", createBlog)
	// router.Put("/api/blogs/{id}", updateBlog)
	// router.Delete("/api/blogs/{id}", deleteBlog)
}
