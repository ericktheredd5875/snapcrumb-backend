CREATE TABLE url_visits (
    id SERIAL PRIMARY KEY,
    short_code TEXT NOT NULL,
    visited_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    ip_address TEXT,
    user_agent TEXT
);