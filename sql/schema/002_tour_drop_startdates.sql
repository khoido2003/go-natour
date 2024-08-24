-- +goose Up
ALTER TABLE tours DROP COLUMN startDates;

-- +goose Down
ALTER TABLE tours ADD COLUMN startDates TIMESTAMP[];
