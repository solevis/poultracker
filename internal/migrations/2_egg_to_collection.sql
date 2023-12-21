-- +migrate Up
ALTER TABLE egg RENAME TO collection;

-- +migrate Down
ALTER TABLE collection RENAME TO egg;
