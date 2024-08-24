package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (apiConfig *APIConfig) GetAllTours(w http.ResponseWriter, r *http.Request) {
	tours, err := apiConfig.DB.GetAllTours(r.Context())

	fmt.Println(tours)
	if err != nil {
		fmt.Println(err)
		respondWithError(w, http.StatusInternalServerError, "Failed to fetch tours")
		return
	}

	respondWithJson(w, http.StatusOK, tours)
}

/////////////////////////////////////////////////

// Create new tour
func (apiConfig *APIConfig) CreateTour(w http.ResponseWriter, r *http.Request) {

	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}

	err := decoder.Decode(&params)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("TEST: ", params.Name)

	respondWithJson(w, http.StatusAccepted, "")
}
