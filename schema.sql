CREATE DATABASE watch_party;

CREATE TABLE users (
id SERIAL PRIMARY KEY,
name VARCHAR(100),
email VARCHAR(100));

CREATE TABLE events (
id SERIAL PRIMARY KEY,
title VARCHAR(100),
type VARCHAR(100),
city VARCHAR(100),
start_time TIMESTAMP);

CREATE TABLE venues (
id SERIAL PRIMARY KEY,
city VARCHAR(100),
address VARCHAR(150),
description VARCHAR(300));

CREATE TABLE watchers (
event_id REFERENCES events,
user_id REFERENCES users);

CREATE TABLE hosters (
event_id REFERENCES events,
venue_id REFERENCES venues);

CREATE TABLE interests (
type VARCHAR(100),
city VARCHAR(100),
user_id REFERENCES users);