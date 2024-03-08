CREATE TABLE questions (
	id uuid NOT NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	questioning varchar(255) NULL,
	"type" varchar(255) NULL,
	image varchar(255) NULL,
	alternatives _text NULL,
	answer varchar(255) NULL,
	discipline varchar(255) NULL,
	CONSTRAINT questions_pkey PRIMARY KEY (id)
);
CREATE INDEX idx_questions_deleted_at ON questions USING btree (deleted_at);