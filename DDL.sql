-- create database bike_rent_express
CREATE DATABASE bike_rent_express;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE user_role AS ENUM ('ADMIN', 'USER');
CREATE TYPE vehicle_status AS ENUM ('AVAILABLE', 'NOT_AVAILABLE');

-- tabel user
CREATE TABLE users(
	id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
	name VARCHAR(255) NOT NULL,
	username VARCHAR(255) NOT NULL,
	password VARCHAR(255) NOT NULL,
	address VARCHAR(255) NULL,
	role user_role NOT NULL,
	can_rent BOOLEAN NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT NULL,
	telp VARCHAR(255) NOT NULL,
	deleted_at DATE NULL
);

-- tabel balance
CREATE TABLE balance(
	id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
	amount INTEGER NOT NULL,
	user_id uuid NOT NULL REFERENCES users(id) ON UPDATE CASCADE ON DELETE CASCADE,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- tabel employee
CREATE TABLE employee(
	id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
	name VARCHAR(255) NOT NULL,
	telp VARCHAR(255) NOT NULL,
	username VARCHAR(255) NOT NULL,
	password VARCHAR(255) NOT NULL,
	deleted_at DATE NULL
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- tabel motor_vehicle
CREATE TABLE motor_vehicle(
	id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
	name VARCHAR(255) NOT NULL,
	type VARCHAR(255) NOT NULL,
	price INTEGER NOT NULL,
	plat VARCHAR(255) NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	production_year VARCHAR(255) NOT NULL,
	status vehicle_status NOT NULL,
	deleted_at DATE NULL
);

-- tabel transaction
CREATE TABLE transaction(
	id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
	user_id uuid NOT NULL REFERENCES users(id),
	motor_vehicle_id uuid NOT NULL REFERENCES motor_vehicle(id),
	start_date DATE NOT NULL,
	end_date DATE NOT NULL,
	price INTEGER NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	employee_id uuid NOT NULL REFERENCES employee(id)
);

-- tabel motor_return
CREATE TABLE motor_return(
	id uuid DEFAULT uuid_generate_V4() PRIMARY KEY,
	transaction_id uuid NOT NULL REFERENCES transaction(id),
	return_date DATE NOT NULL,
	extra_charge INTEGER NOT NULL,
	condition_motor VARCHAR(255) NOT NULL,
	description VARCHAR(255) NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

