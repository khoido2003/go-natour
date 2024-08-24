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
  imageCover
)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
RETURNING *;

-- name: GetAllTours :many
SELECT * FROM tours; 