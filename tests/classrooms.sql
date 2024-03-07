CREATE TABLE classrooms (
	id uuid NOT NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	"name" varchar(255) NULL,
	description varchar(255) NULL,
	course varchar(255) NULL,
	CONSTRAINT classrooms_pkey PRIMARY KEY (id),
	CONSTRAINT uni_classrooms_name UNIQUE (name)
);
CREATE INDEX idx_classrooms_deleted_at ON classrooms USING btree (deleted_at);