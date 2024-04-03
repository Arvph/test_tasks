-- +goose Up
CREATE TABLE "task" (
    "id" SERIAL PRIMARY KEY,
    "user_id" INTEGER NOT NULL,
    "text" VARCHAR(255),
    "created_at" TIMESTAMP DEFAULT now(),
    "status" VARCHAR(10) NOT NULL
);



-- +goose Down
DROP TABLE task;
-- DROP TABLE user;

