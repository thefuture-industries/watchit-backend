-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS user_cores (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    user_uuid VARCHAR(255) NOT NULL UNIQUE,
    email VARCHAR(100) NULL UNIQUE,
    verified BOOLEAN DEFAULT FALSE
);

CREATE OR REPLACE FUNCTION set_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_updated_at_trigger
    BEFORE UPDATE ON user_cores
    FOR EACH ROW
    EXECUTE FUNCTION set_updated_at();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TRIGGER IF EXISTS set_updated_at_trigger ON user_cores;
DROP FUNCTION IF EXISTS set_updated_at();
DROP TABLE IF EXISTS user_cores;
-- +goose StatementEnd
