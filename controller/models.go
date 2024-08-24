package controller

import (
	"time"

	"github.com/google/uuid"
	"github.com/khoido2003/go-natour/database"
)

type Tour struct {
	ID             uuid.UUID `json:"id"`
	Name           string    `json:"name"`
	Slug           string    `json:"slug"`
	Duration       float64   `json:"duration"`
	Maxgroupsize   int32     `json:"max_group_size"`
	Difficulty     string    `json:"difficulty"`
	Ratingaverage  float64   `json:"rating_average"`
	Ratingquantity int32     `json:"rating_quantity"`
	Price          int32     `json:"price"`
	Discount       float64   `json:"discount"`
	Summary        string    `json:"summary"`
	Description    string    `json:"description"`
	Imagecover     string    `json:"image_cover"`
	Createdat      time.Time `json:"created_at"`
	// Startdates     []time.Time `json:"start_dates"`
}

/////////////////////////////////////////////

func dbTourToTour(tour database.Tour) Tour {

	return Tour{
		ID:             tour.ID,
		Name:           tour.Name,
		Slug:           tour.Slug.String,
		Duration:       tour.Duration.Float64,
		Maxgroupsize:   tour.Maxgroupsize,
		Difficulty:     tour.Difficulty,
		Ratingaverage:  tour.Ratingaverage.Float64,
		Ratingquantity: tour.Ratingquantity.Int32,
		Price:          tour.Price,
		Discount:       tour.Discount.Float64,
		Summary:        tour.Summary,
		Description:    tour.Description.String,
		Imagecover:     tour.Imagecover.String,
		Createdat:      tour.Createdat,
	}
}

func dbToursToTours(tours []database.Tour) []Tour {
	tourList := make([]Tour, len(tours))
	for i, tour := range tours {
		tourList[i] = dbTourToTour(tour)
	}
	return tourList
}
