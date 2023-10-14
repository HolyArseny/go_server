package database

type Schema struct {
	name  string
	query string
}

var Schemas = []Schema{
	Schema{
		"users",
		`CREATE TABLE IF NOT EXISTS "users" (
		  "id" BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
		  "email" VARCHAR(255) NOT NULL UNIQUE,
		  "password" VARCHAR NOT NULL,
		  "name" VARCHAR(25) NOT NULL UNIQUE,
		  "profile_image" VARCHAR,
		  "inserted_at" TIMESTAMP DEFAULT current_timestamp,
		  "updated_at" TIMESTAMP DEFAULT current_timestamp
		);`,
	},
	Schema{
		"user_session",
		`CREATE TABLE IF NOT EXISTS "user_session" (
		  "sid" VARCHAR NOT NULL COLLATE "default",
		  "sess" JSON NOT NULL,
		  "expire" timestamp(6) NOT NULL
		);
		WITH (OIDS=FALSE);
		ALTER TABLE "user_session" ADD CONSTRAINT "user_session_pkey" PRIMARY KEY ("sid") NOT DEFERRABLE INITIALLY IMMEDIATE;
		CREATE INDEX "IDX_user_session_expire" ON "user_session" ("expire");`,
	},
	Schema{
		"staff_roles",
		`CREATE TABLE IF NOT EXISTS "staff_roles" (
		  "name" VARCHAR(255) NOT NULL UNIQUE PRIMARY KEY,
		  "inserted_at" TIMESTAMP DEFAULT current_timestamp,
		  "updated_at" TIMESTAMP DEFAULT current_timestamp
		);`,
	},
	Schema{
		"staff",
		`CREATE TABLE IF NOT EXISTS "staff" (
		  "id" BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
		  "email" VARCHAR(255) NOT NULL UNIQUE,
		  "password" VARCHAR NOT NULL,
		  "name" VARCHAR(25) NOT NULL UNIQUE,
		  "profile_image" VARCHAR,
		  "role" VARCHAR(255) DEFAULT 'unknown',
		  "inserted_at" TIMESTAMP DEFAULT current_timestamp,
		  "updated_at" TIMESTAMP DEFAULT current_timestamp,
		  CONSTRAINT fk_role FOREIGN KEY(role) REFERENCES staff_roles(name)
		);`,
	},
	Schema{
		"staff_session", `
		CREATE TABLE IF NOT EXISTS "staff_session" (
		  "sid" VARCHAR NOT NULL COLLATE "default",
		  "sess" JSON NOT NULL,
		  "expire" timestamp(6) NOT NULL
		);
		WITH (OIDS=FALSE);
		ALTER TABLE "staff_session" ADD CONSTRAINT "staff_session_pkey" PRIMARY KEY ("sid") NOT DEFERRABLE INITIALLY IMMEDIATE;
		CREATE INDEX "IDX_staff_session_expire" ON "staff_session" ("expire");,`,
	},
	Schema{
		"subscriptions_type",
		`DO $$
		BEGIN
		    IF NOT EXISTS (
		        SELECT 1
		        FROM pg_type
		        WHERE typname = 'your_type_name'
		    ) THEN
   				CREATE TYPE IF NOT EXISTS subscription_type_recurrency AS ENUM
				  ('month', 'one time');
		    END IF;
		END $$;
		CREATE TABLE IF NOT EXISTS "subscriptions_type" (
		  "id" BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
		  "name" VARCHAR,
		  "amount" INT,
		  "importance" INT DEFAULT 0,
		  "active" BOOLEAN default FALSE,
		  "recurrency" subscription_type_recurrency,
		  "options" JSON DEFAULT ('[]'),
		  "inserted_at" TIMESTAMP DEFAULT current_timestamp,
		  "updated_at" TIMESTAMP DEFAULT current_timestamp
		);`,
	},
	Schema{
		"subscriptions",
		`DO $$
		BEGIN
		    IF NOT EXISTS (
		        SELECT 1
		        FROM pg_type
		        WHERE typname = 'your_type_name'
		    ) THEN
      			CREATE TYPE subscription_status AS ENUM
				  ('active', 'past_due', 'canceled', 'incomplete');
		    END IF;
		END $$;
		CREATE TABLE "subscriptions" (
		  "id" BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
		  "user_id" BIGINT NOT NULL,
		  "status" subscription_status,
		  "type_id" BIGINT NOT NULL,
		  "active_due" TIMESTAMP,
		  "start_from" TIMESTAMP,
		  "inserted_at" TIMESTAMP DEFAULT current_timestamp,
		  "updated_at" TIMESTAMP DEFAULT current_timestamp,
		  CONSTRAINT fk_type FOREIGN KEY(type_id) REFERENCES subscriptions_type(id),
		  CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES users(id)
		);`,
	},
	Schema{
		"courses_groups",
		`CREATE TABLE IF NOT EXISTS "courses_groups" (
		  "id" BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
		  "name" VARCHAR(255) NOT NULL,
		  "inserted_at" TIMESTAMP DEFAULT current_timestamp,
		  "updated_at" TIMESTAMP DEFAULT current_timestamp
		);`,
	},
	Schema{
		"courses",
		`CREATE TABLE IF NOT EXISTS "courses" (
		  "id" BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
		  "name" VARCHAR(255) NOT NULL,
		  "difficulty" FLOAT NOT NULL,
		  "type" VARCHAR DEFAULT 'common' NOT NULL,
		  "course_group" BIGINT NOT NULL,
		  "inserted_at" TIMESTAMP DEFAULT current_timestamp,
		  "updated_at" TIMESTAMP DEFAULT current_timestamp,
		  CONSTRAINT fk_group FOREIGN KEY(course_group) REFERENCES courses_groups(id)
		);`,
	},
	Schema{
		"course_topic",
		`CREATE TABLE IF NOT EXISTS "course_topic" (
		  "id" BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
		  "name" VARCHAR NOT NULL,
		  "active" BOOLEAN DEFAULT FALSE,
		  "course" BIGINT NOT NULL,
		  "difficulty" FLOAT NOT NULL,
		  "inserted_at" TIMESTAMP DEFAULT current_timestamp,
		  "updated_at" TIMESTAMP DEFAULT current_timestamp,
		  CONSTRAINT fk_course FOREIGN KEY(course) REFERENCES courses(id)
		);`,
	},
	Schema{
		"chaqllenges",
		`CREATE TABLE IF NOT EXISTS "challenges" (
		  "id" BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
		  "data" JSON NOT NULL,
		  "difficulty" FLOAT NOT NULL,
		  "scores" INT DEFAULT 0,
		  "course_topic" BIGINT NOT NULL,
		  "inserted_at" TIMESTAMP DEFAULT current_timestamp,
		  "updated_at" TIMESTAMP DEFAULT current_timestamp,
		  CONSTRAINT fk_course_topic FOREIGN KEY(course_topic) REFERENCES course_topic(id)
		);`,
	},
	Schema{
		"theory",
		`CREATE TABLE IF NOT EXISTS "theory" (
		  "id" BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
		  "name" VARCHAR NOT NULL,
		  "data" JSON NOT NULL,
		  "course_topic" BIGINT NOT NULL,
		  "difficulty" FLOAT NOT NULL,
		  "inserted_at" TIMESTAMP DEFAULT current_timestamp,
		  "updated_at" TIMESTAMP DEFAULT current_timestamp,
		  CONSTRAINT fk_course_topic FOREIGN KEY(course_topic) REFERENCES course_topic(id)
		);`,
	},
	Schema{
		"user_progress",
		`CREATE TABLE IF NOT EXISTS "users_progress" (
		  "id" BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
		  "user_id" BIGINT NOT NULL,
		  "challenge_id" BIGINT NOT NULL,
		  "score" INT DEFAULT 0,
		  "inserted_at" TIMESTAMP DEFAULT current_timestamp,
		  "updated_at" TIMESTAMP DEFAULT current_timestamp,
		  CONSTRAINT fk_challenge_id FOREIGN KEY(challenge_id) REFERENCES challenges(id),
		  CONSTRAINT fk_user_id FOREIGN KEY(user_id) REFERENCES users(id)
		);`,
	},
}
