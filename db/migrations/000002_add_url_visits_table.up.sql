CREATE TABLE url_visits (
    "id" SERIAL NOT NULL,
	"shortcode" VARCHAR(10) NOT NULL DEFAULT '',
	"visited_at" TIMESTAMPTZ NOT NULL DEFAULT now(),
	"ip_address" VARCHAR(100) NOT NULL DEFAULT '',
	"user_agent" VARCHAR(250) NOT NULL DEFAULT '',
	"referer" VARCHAR(150) NOT NULL DEFAULT '',
	PRIMARY KEY ("id")
);

CREATE INDEX IF NOT EXISTS shortcode ON url_visits (shortcode);



ALTER TABLE "url_visits"
	ALTER COLUMN "user_agent" TYPE VARCHAR(250),
	ALTER COLUMN "user_agent" SET NOT NULL,
	ALTER COLUMN "user_agent" SET DEFAULT '';