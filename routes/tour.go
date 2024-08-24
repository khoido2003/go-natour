package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/khoido2003/go-natour/controller"
)

func tourRoute(router *chi.Mux, apiConfig *controller.APIConfig) {

	router.Get("/tours", apiConfig.GetAllTours)
	router.Post("/tours", apiConfig.CreateTour)

}
