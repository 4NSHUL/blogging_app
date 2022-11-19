package main

import "github.com/go-chi/chi"

func RegisterRoutes(router *chi.Mux) {
	router.Get("/api/blogs", getBlogs)
	router.Get("/api/blogs/{id}", getBlog)
	router.Post("/api/blogs/create", createBlog)
	router.Put("/api/blogs/update/{id}", updateBlog)
	router.Delete("/api/blogs/delete/{id}", deleteBlog)
}
