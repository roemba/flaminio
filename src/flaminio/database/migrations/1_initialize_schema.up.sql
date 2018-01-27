CREATE TABLE IF NOT EXISTS flaminio.users (
uuid uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
createdAt timestamp NOT NULL DEFAULT (NOW() AT TIME ZONE 'utc'),
updatedAt timestamp NOT NULL DEFAULT (NOW() AT TIME ZONE 'utc'),
firstName character varying(255) NOT NULL,
middleName character varying(255),
lastName character varying(255) NOT NULL,
password bytea NOT NULL,
email citext NOT NULL UNIQUE,
preferredLocale character varying(10) NOT NULL DEFAULT 'en');

INSERT INTO flaminio.users (firstName, lastName, password, email) VALUES ('admin', 'admin', '$2a$14$.0PvyKvM6jRcOPdOFaozjeu5PUttosT9BT5pmGxINbjbA4gHqIstK', 'admin@admin.com');