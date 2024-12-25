package models

var Courses = []Schema{
	{
		0,
		"courses_groups",
		`CREATE TABLE "courses_groups" (
		    "id" BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
		    "name" VARCHAR(255) NOT NULL,
		    "inserted_at" TIMESTAMP DEFAULT current_timestamp,
		    "updated_at" TIMESTAMP DEFAULT current_timestamp
		  );
		`,
	},
	{
		1,
		"courses",
		`CREATE TABLE "courses" (
		    "id" BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
		    "name" VARCHAR(255) NOT NULL,
		    "difficulty" FLOAT NOT NULL,
		    "type" VARCHAR DEFAULT 'common' NOT NULL,
		    "course_group" BIGINT NOT NULL,
		    "inserted_at" TIMESTAMP DEFAULT current_timestamp,
		    "updated_at" TIMESTAMP DEFAULT current_timestamp,
		    CONSTRAINT fk_group FOREIGN KEY(course_group) REFERENCES courses_groups(id)
		  );
		`,
	},
	{
		2,
		"course_topic",
		`CREATE TABLE "course_topic" (
		    "id" BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
		    "name" VARCHAR NOT NULL,
		    "active" BOOLEAN DEFAULT FALSE,
		    "course" BIGINT NOT NULL,
		    "difficulty" FLOAT NOT NULL,
		    "inserted_at" TIMESTAMP DEFAULT current_timestamp,
		    "updated_at" TIMESTAMP DEFAULT current_timestamp,
		    CONSTRAINT fk_course FOREIGN KEY(course) REFERENCES courses(id)
		  );
		`,
	},
}
