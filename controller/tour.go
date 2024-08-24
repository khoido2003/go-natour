package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/khoido2003/go-natour/database"
)

// Get all tours
func (apiConfig *APIConfig) GetAllTours(w http.ResponseWriter, r *http.Request) {
	tours, err := apiConfig.DB.GetAllTours(r.Context())

	if err != nil {
		fmt.Println(err)
		respondWithError(w, http.StatusInternalServerError, "Failed to fetch tours")
		return
	}

	respondWithJson(w, http.StatusOK, dbToursToTours(tours))
}

/////////////////////////////////////////////////

// Create new tour
func (apiConfig *APIConfig) CreateTour(w http.ResponseWriter, r *http.Request) {

	type parameters struct {
		Name           string  `json:"name"`
		Duration       float64 `json:"duration"`
		Maxgroupsize   int32   `json:"max_group_size"`
		Difficulty     string  `json:"difficulty"`
		Ratingaverage  float64 `json:"rating_average"`
		Ratingquantity int32   `json:"rating_quantity"`
		Price          int32   `json:"price"`
		Summary        string  `json:"summary"`
		Description    string  `json:"description"`
		Imagecover     string  `json:"image_cover"`
	}

	decoder := json.NewDecoder(r.Body)

	payload := parameters{}

	err := decoder.Decode(&payload)

	if err != nil {
		fmt.Println(err)
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	tour, err := apiConfig.DB.CreateTour(r.Context(), database.CreateTourParams{
		Name:           payload.Name,
		Duration:       sql.NullFloat64{Float64: payload.Duration, Valid: payload.Duration != 0},
		Maxgroupsize:   payload.Maxgroupsize,
		Difficulty:     payload.Difficulty,
		Ratingaverage:  sql.NullFloat64{Float64: payload.Ratingaverage, Valid: payload.Ratingaverage != 0},
		Ratingquantity: sql.NullInt32{Int32: payload.Ratingquantity, Valid: payload.Ratingquantity != 0},
		Price:          payload.Price,
		Summary:        payload.Summary,
		Description:    sql.NullString{String: payload.Description, Valid: payload.Description != ""},
		Imagecover:     sql.NullString{String: payload.Imagecover, Valid: payload.Imagecover != ""},
	})

	if err != nil {
		fmt.Println(err)
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	respondWithJson(w, http.StatusAccepted, dbTourToTour(tour))
}
