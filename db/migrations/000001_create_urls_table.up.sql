CREATE TABLE urls (
  id integer NOT NULL DEFAULT nextval('urls_id_seq'::regclass),
  original_url text COLLATE pg_catalog."default" NOT NULL,
  shortcode character varying(10) COLLATE pg_catalog."default" NOT NULL,
  created_at timestamp without time zone DEFAULT now(),
  expires_at timestamp without time zone,
  visit_count integer DEFAULT 0,
  CONSTRAINT urls_pkey PRIMARY KEY (id),
  CONSTRAINT urls_shortcode_key UNIQUE (shortcode)
);