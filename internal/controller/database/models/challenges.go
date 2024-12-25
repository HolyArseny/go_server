package models

var Challenges = []Schema{
	{
		3,
		"challenges",
		`CREATE TABLE "challenges" (
		    "id" BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
		    "data" JSON NOT NULL,
		    "difficulty" FLOAT NOT NULL,
		    "scores" INT DEFAULT 0,
		    "course_topic" BIGINT NOT NULL,
		    "inserted_at" TIMESTAMP DEFAULT current_timestamp,
		    "updated_at" TIMESTAMP DEFAULT current_timestamp,
		    CONSTRAINT fk_course_topic FOREIGN KEY(course_topic) REFERENCES course_topic(id)
		  );
		`,
	},
}
