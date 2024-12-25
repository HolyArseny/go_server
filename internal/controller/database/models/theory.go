package models

var Theory = []Schema{
	{
		3,
		"theory",
		`CREATE TABLE "theory" (
		    "id" BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
		    "data" VARCHAR NOT NULL,
		    "course_topic" BIGINT NOT NULL,
		    "inserted_at" TIMESTAMP DEFAULT current_timestamp,
		    "updated_at" TIMESTAMP DEFAULT current_timestamp,
		    CONSTRAINT fk_course_topic FOREIGN KEY(course_topic) REFERENCES course_topic(id)
		  );
		`,
	},
}
