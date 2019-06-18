-- +migrate Up
ALTER TABLE employee ALTER COLUMN mobile TYPE varchar;

-- +migrate Down
