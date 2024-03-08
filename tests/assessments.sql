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

CREATE TABLE assessments (
	id uuid NOT NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	description varchar(22) NULL,
	courses _text NULL,
	classrooms _text NULL,
	start_date date NULL,
	end_date date NULL,
	quiz text NULL,
	CONSTRAINT assessments_pkey PRIMARY KEY (id)
);
CREATE INDEX idx_assessments_deleted_at ON assessments USING btree (deleted_at);