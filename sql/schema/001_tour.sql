-- +goose Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE tours (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v1(),
  name VARCHAR(100) NOT NULL UNIQUE,
  slug VARCHAR(255), 
  duration DOUBLE PRECISION,
  maxGroupSize INTEGER NOT NULL,
  difficulty VARCHAR(255) NOT NULL CHECK (difficulty IN ('easy', 'medium', 'difficult')),
  ratingAverage DOUBLE PRECISION DEFAULT 4.5 CHECK (ratingAverage >= 1),
  ratingQuantity INTEGER DEFAULT 0,
  price INTEGER NOT NULL,
  discount DOUBLE PRECISION,
  summary TEXT NOT NULL,
  description TEXT,
  imageCover TEXT,
  createdAt TIMESTAMP NOT NULL DEFAULT NOW(),
  startDates TIMESTAMP[],
  secretTour BOOLEAN DEFAULT FALSE
);


-- +goose Down
DROP TABLE tours;