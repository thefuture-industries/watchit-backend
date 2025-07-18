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

INSERT INTO genres (genre_id, genre_name) VALUES
(18, 'Drama'),
(36, 'History'),
(27, 'Horror'),
(16, 'Animation'),
(99, 'Documentary'),
(10751, 'Family'),
(10402, 'Music'),
(9648, 'Mystery'),
(878, 'Science Fiction'),
(28, 'Action'),
(35, 'Comedy'),
(80, 'Crime'),
(14, 'Fantasy'),
(37, 'Western'),
(12, 'Adventure'),
(10749, 'Romance'),
(10770, 'TV Movie'),
(53, 'Thriller'),
(10752, 'War');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TRIGGER IF EXISTS set_updated_at_trigger ON genres;
DROP FUNCTION IF EXISTS set_updated_at();
DROP TABLE IF EXISTS genres;
-- +goose StatementEnd
