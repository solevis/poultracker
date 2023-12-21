-- +migrate Up
CREATE TABLE IF NOT EXISTS "egg" (
    "id" INTEGER NOT NULL PRIMARY KEY, 
    "laid_date" DATE NOT NULL UNIQUE, 
    "number" INTEGER NOT NULL
);

-- +migrate Down
DROP TABLE "egg";
