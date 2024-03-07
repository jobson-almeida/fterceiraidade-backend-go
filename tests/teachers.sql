CREATE TABLE teachers (
	id uuid NOT NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	avatar varchar(255) NULL,
	fistname varchar(255) NULL,
	lastname varchar(255) NULL,
	email varchar(255) NULL,
	phone varchar(255) NULL,
	address bytea NULL,
	firstname varchar(255) NULL,
	CONSTRAINT idx_teachers_email UNIQUE (email),
	CONSTRAINT teachers_pkey PRIMARY KEY (id)
);
CREATE INDEX idx_teachers_deleted_at ON teachers USING btree (deleted_at);