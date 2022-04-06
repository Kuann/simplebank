
CREATE TABLE account (
	id serial PRIMARY KEY,
	username VARCHAR ( 50 ) UNIQUE NOT NULL,
	password VARCHAR ( 50 ) NOT NULL
);



CREATE TABLE employee (
	id serial PRIMARY KEY,
	employee_name VARCHAR ( 50 ) UNIQUE NOT NULL
);
