CREATE TABLE students (
	id uuid NOT NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	avatar varchar(255) NULL,
	firstname varchar(255) NULL,
	lastname varchar(255) NULL,
	email varchar(255) NULL,
	phone varchar(255) NULL,
	address bytea NULL,
	CONSTRAINT idx_students_email UNIQUE (email),
	CONSTRAINT students_pkey PRIMARY KEY (id)
);
CREATE INDEX idx_students_deleted_at ON students USING btree (deleted_at);
 