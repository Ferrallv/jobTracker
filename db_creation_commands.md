# Command List

Enter these into your Postgres CLI to create the database for the app. Be sure to change the values between '<>'.

	CREATE DATABASE jobTracker;

	\c jobTracker

	CREATE USER <username> WITH PASSWORD <password>;
	
	GRANT ALL PRIVILEGES ON DATABASE jobtracker TO <username>;

	GRANT USAGE, SELECT ON ALL SEQUENCES IN SCHEMA public TO <username>;
	
	CREATE EXTENSION citext;
		
	CREATE TABLE contacts (
		id 			serial PRIMARY KEY,
		name 		varchar(45) NOT NULL, position varchar(45), number varchar(15),
		email 		citext,
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
		id 		serial PRIMARY KEY,
		date 	bigint NOT NULL,
		method 	varchar(10) NOT NULL,
		job_id 	int REFERENCES application(id)
	);