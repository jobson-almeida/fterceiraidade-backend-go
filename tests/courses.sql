CREATE TABLE courses (
	id uuid NOT NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	"name" varchar(255) NULL,
	description varchar(255) NULL,
	image varchar(255) NULL,
	CONSTRAINT courses_pkey PRIMARY KEY (id),
	CONSTRAINT uni_courses_name UNIQUE (name)
);
CREATE INDEX idx_courses_deleted_at ON courses USING btree (deleted_at);