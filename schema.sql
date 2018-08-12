CREATE DATABASE watch_party;

CREATE TABLE users (
id SERIAL,
name VARCHAR(100),
email VARCHAR(100)));

CREATE TABLE events (
id SERIAL,
title VARCHAR(100),
type VARCHAR(100),
city VARCHAR(100),
start_time DATE);

CREATE TABLE venues (
id SERIAL,
city VARCHAR(100),
address VARCHAR(150),
description VARCHAR(300));

CREATE TABLE watchers (
event_id INT,
user_id INT);

CREATE TABLE hosters (
event_id INT,
venue_id INT);

CREATE TABLE interests (
type VARCHAR(100),
city VARCHAR(100),
user_id INT);