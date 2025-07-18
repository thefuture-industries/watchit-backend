-- +goose Up
-- +goose StatementBegin
CREATE TABLE movie_genres (
    movie_id BIGINT REFERENCES movie(id) ON DELETE CASCADE,
    genre_id INTEGER,

    PRIMARY KEY (movie_id, genre_id)
);


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS movie_genres;
-- +goose StatementEnd
