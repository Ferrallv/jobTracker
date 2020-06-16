# Command List

Enter these into your Postgres CLI to create the database for the app. Be sure to change the values between '<>'.

	CREATE DATABASE jobTracker;

	\c jobTracker

	CREATE USER <username> WITH PASSWORD '<password>';
	
	CREATE EXTENSION citext;

	CREATE DOMAIN email AS citext
  CHECK ( value ~ '^[a-zA-Z0-9.!#$%&''*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$' );
		
	CREATE TABLE contacts (
		id 			serial PRIMARY KEY,
		name 		varchar(45) NOT NULL,
		position 	varchar(45),
		number 		varchar(45),
		email 		email,
		company 	varchar(45) NOT NULL,
		note 		text
	); 

	CREATE TABLE application (
		id serial 	PRIMARY KEY,
		job_title 	varchar(45) NOT NULL,
		description text,
		url 		text,
		company 	varchar(45) NOT NULL,
		resume 		bytea,
		cvr_letter 	bytea,
		app_date 	bigint,
		offer 		bigint,
		rejected 	bigint,
		declined 	bigint
	);

	CREATE TABLE interview (
		id 			serial PRIMARY KEY,
		date 		bigint NOT NULL,
		method 		varchar(45) NOT NULL,
		job_id 		int REFERENCES application(id)
	);

	GRANT ALL PRIVILEGES ON DATABASE jobtracker TO <username>;

	GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO <username>;

	GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public TO <username>;

	ALTER DEFAULT PRIVILEGES FOR USER <username> IN SCHEMA public GRANT ALL PRIVILEGES ON TABLES TO <username>;