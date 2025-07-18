-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS limits (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    index_limit INTEGER NOT NULL UNIQUE,
    max_query_length INTEGER NOT NULL DEFAULT 255 CHECK (max_query_length IN (255, 1000)), -- Длина поискового запроса (макс)
    daily_search_limit INTEGER NOT NULL DEFAULT 10 CHECK (daily_search_limit IN (10, 100)), -- Кол-во поисков в день
    search_priority INTEGER NOT NULL DEFAULT 0 CHECK (search_priority IN (0, 1)) -- Приоритет в очереди (0 — низкий)
);

CREATE OR REPLACE FUNCTION set_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_updated_at_trigger
    BEFORE UPDATE ON limits
    FOR EACH ROW
    EXECUTE FUNCTION set_updated_at();

INSERT INTO limits (index_limit, max_query_length, daily_search_limit, search_priority) VALUES (101010, 255, 10, 0);
INSERT INTO limits (index_limit, max_query_length, daily_search_limit, search_priority) VALUES (111111, 1000, 100, 1);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TRIGGER IF EXISTS set_updated_at_trigger ON limits;
DROP FUNCTION IF EXISTS set_updated_at();
DROP TABLE IF EXISTS limits;
-- +goose StatementEnd
