BEGIN;

CREATE TABLE companies (
	id serial NOT NULL PRIMARY KEY,
	created_at timestamptz NOT NULL,
	updated_at timestamptz NOT NULL,
	name text NOT NULL,
	address text NOT NULL
);


CREATE TABLE users (
	id serial NOT NULL PRIMARY KEY,
	created_at timestamptz NOT NULL,
	updated_at timestamptz NOT NULL,
	email text NOT NULL,
	salt text NOT NULL,
	passhash text NOT NULL,
	name text NOT NULL,
	phone text NOT NULL
);


COMMIT;
