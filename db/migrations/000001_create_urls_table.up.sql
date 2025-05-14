CREATE TABLE urls (
  "id" SERIAL NOT NULL,
	"original_url" TEXT NOT NULL,
	"shortcode" VARCHAR(10) NOT NULL,
	"created_at" TIMESTAMP NULL DEFAULT now(),
	"expires_at" TIMESTAMP NULL DEFAULT NULL,
	"visit_count" INTEGER NULL DEFAULT 0,
	PRIMARY KEY ("id"),
	UNIQUE ("shortcode")
);