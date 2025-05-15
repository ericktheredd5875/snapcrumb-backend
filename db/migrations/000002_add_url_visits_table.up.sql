CREATE TABLE url_visits (
    "id" SERIAL NOT NULL,
	"shortcode" VARCHAR(10) NOT NULL DEFAULT '',
	"visited_at" TIMESTAMPTZ NOT NULL DEFAULT now(),
	"ip_address" VARCHAR(100) NOT NULL DEFAULT '',
	"user_agent" VARCHAR(150) NOT NULL DEFAULT '',
	PRIMARY KEY ("id")
);

CREATE INDEX IF NOT EXISTS shortcode ON url_visits (shortcode);