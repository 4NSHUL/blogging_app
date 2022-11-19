package routing

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

var router *chi.Mux

func init() {
	router = chi.NewRouter()
	configureRouter()
}

func configureRouter() {
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.StripSlashes)
}

func GetRouter() *chi.Mux {
	return router
}
