-- name: CreateTour :one

INSERT INTO tours (
  name,
  duration,
  maxGroupSize,
  difficulty,  
  ratingAverage,
  ratingQuantity,
  price,
  summary,
  description,
  imageCover,
  startDates
)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
RETURNING *;

-- name: GetAllTours :many
SELECT * FROM tours; 