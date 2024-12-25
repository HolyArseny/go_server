package models

var Users = []Schema{
	{
		0,
		"users",
		`CREATE TABLE "users" (
		  "id" BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
		  "email" VARCHAR(255) NOT NULL UNIQUE,
		  "password" VARCHAR NOT NULL,
		  "name" VARCHAR(25) NOT NULL UNIQUE,
		  "profile_image" VARCHAR,
		  "inserted_at" TIMESTAMP DEFAULT current_timestamp,
		  "updated_at" TIMESTAMP DEFAULT current_timestamp
		);`,
	},
	// {
	// 	0,
	// 	"user_session",
	// 	`CREATE TABLE "user_session" (
	// 	    "sid" VARCHAR NOT NULL COLLATE "default",
	// 	    "sess" JSON NOT NULL,
	// 	    "expire" timestamp(6) NOT NULL,
	// 		CONSTRAINT "user_session_pkey" PRIMARY KEY ("sid") NOT DEFERRABLE INITIALLY IMMEDIATE
	// 	  ) WITH (OIDS=FALSE);
	// 	  CREATE INDEX "IDX_user_session_expire" ON "user_session" ("expire");
	// 	`,
	// },
	{
		4,
		"users_progress",
		`CREATE TABLE "users_progress" (
		    "id" BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
		    "user_id" BIGINT NOT NULL,
		    "challenge_id" BIGINT NOT NULL,
		    "score" INT DEFAULT 0,
		    "inserted_at" TIMESTAMP DEFAULT current_timestamp,
		    "updated_at" TIMESTAMP DEFAULT current_timestamp,
		    CONSTRAINT fk_challenge_id FOREIGN KEY(challenge_id) REFERENCES challenges(id),
		    CONSTRAINT fk_user_id FOREIGN KEY(user_id) REFERENCES users(id)
		  );
		`,
	},
}
