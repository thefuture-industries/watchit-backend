-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS payment_cards (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    user_uuid VARCHAR(255) NOT NULL,
    card_number VARCHAR(16) NOT NULL,
    expiration_month SMALLINT NOT NULL CHECK(expiration_month BETWEEN 1 AND 12), -- month card
    expiration_year SMALLINT NOT NULL -- year card
);

CREATE OR REPLACE FUNCTION set_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_updated_at_trigger
    BEFORE UPDATE ON payment_cards
    FOR EACH ROW
    EXECUTE FUNCTION set_updated_at();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TRIGGER IF EXISTS set_updated_at_trigger ON payment_cards;
DROP FUNCTION IF EXISTS set_updated_at();
DROP TABLE IF EXISTS payment_cards;
-- +goose StatementEnd
