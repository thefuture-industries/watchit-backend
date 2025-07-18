-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOt EXISTS genres (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    genre_id INTEGER NOT NULL UNIQUE,
    genre_name VARCHAR(50) NOT NULL
);

CREATE OR REPLACE FUNCTION set_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_updated_at_trigger
    BEFORE UPDATE ON genres
    FOR EACH ROW
    EXECUTE FUNCTION set_updated_at();

INSERT INTO genres (genre_id, genre_name) VALUES (18, 'Drama');
INSERT INTO genres (genre_id, genre_name) VALUES (36, 'History');
INSERT INTO genres (genre_id, genre_name) VALUES (27, 'Horror');
INSERT INTO genres (genre_id, genre_name) VALUES (16, 'Animation');
INSERT INTO genres (genre_id, genre_name) VALUES (99, 'Documentary');
INSERT INTO genres (genre_id, genre_name) VALUES (10751, 'Family');
INSERT INTO genres (genre_id, genre_name) VALUES (10402, 'Music');
INSERT INTO genres (genre_id, genre_name) VALUES (9648, 'Mystery');
INSERT INTO genres (genre_id, genre_name) VALUES (878, 'Science Fiction');
INSERT INTO genres (genre_id, genre_name) VALUES (28, 'Action');
INSERT INTO genres (genre_id, genre_name) VALUES (35, 'Comedy');
INSERT INTO genres (genre_id, genre_name) VALUES (80, 'Crime');
INSERT INTO genres (genre_id, genre_name) VALUES (14, 'Fantasy');
INSERT INTO genres (genre_id, genre_name) VALUES (37, 'Western');
INSERT INTO genres (genre_id, genre_name) VALUES (12, 'Adventure');
INSERT INTO genres (genre_id, genre_name) VALUES (10749, 'Romance');
INSERT INTO genres (genre_id, genre_name) VALUES (10770, 'TV Movie');
INSERT INTO genres (genre_id, genre_name) VALUES (53, 'Thriller');
INSERT INTO genres (genre_id, genre_name) VALUES (10752, 'War');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TRIGGER IF EXISTS set_updated_at_trigger ON genres;
DROP FUNCTION IF EXISTS set_updated_at();
DROP TABLE IF EXISTS genres;
-- +goose StatementEnd
