CREATE DATABASE IF NOT EXISTS dingtalk;
USE db_name;

CREATE TABLE IF NOT EXISTS atten_list (
	id integer PRIMARY KEY,
	planid integer NOT NULL,
	checktype text NOT NULL,
	approveid integer,
	userid text NOT NULL,
	classid integer NOT NULL,
	classsettingid integer NOT NULL,
	planchecktime text NOT NULL,
	groupid integer NOT NULL,
	createdat text NOT NULL
);

CREATE TABLE IF NOT EXISTS dep_users (
	id integer PRIMARY KEY,
	name text NOT NULL,
	userid text NOT NULL,
	createdat text NOT NULL
);

CREATE TABLE IF NOT EXISTS class_list (
    id integer PRIMARY KEY,
    class_id text NOT NULL,
    class_data text NOT NULL,
    createdat text NOT NULL
);