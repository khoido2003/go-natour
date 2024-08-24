package router

import (
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/khoido2003/go-natour/controller"
)

func RouterController(apiConfig controller.APIConfig) *chi.Mux {

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// logger
	router.Use(middleware.Logger)

	v1Router := chi.NewRouter()

	// Routes
	tourRoute(v1Router, &apiConfig)

	router.Mount("/api/v1", v1Router)

	return router
}
