-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS user_limits (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    user_uuid VARCHAR(255) NOT NULL UNIQUE,
    limit_id INTEGER NOT NULL,
    daily_search_limit_usage INTEGER NOT NULL DEFAULT 0,

    CONSTRAINT fk_limit_id FOREIGN KEY (limit_id) REFERENCES limits(index_limit)
);


CREATE OR REPLACE FUNCTION set_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_updated_at_trigger
    BEFORE UPDATE ON user_limits
    FOR EACH ROW
    EXECUTE FUNCTION set_updated_at();


CREATE OR REPLACE FUNCTION increment_user_limits_usage(p_user_uuid VARCHAR)
RETURNS VOID AS $$
BEGIN
    UPDATE user_limits
    SET
        daily_search_limit_usage = daily_search_limit_usage + 1
    WHERE user_uuid = p_user_uuid;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TRIGGER IF EXISTS set_updated_at_trigger ON user_limits;
DROP FUNCTION IF EXISTS set_updated_at();
DROP TABLE IF EXISTS user_limits;
-- +goose StatementEnd
