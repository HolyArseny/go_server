package models

var Staff = []Schema{
	{
		0,
		"staff_roles",
		`CREATE TABLE "staff_roles" (
		    "name" VARCHAR(255) NOT NULL UNIQUE PRIMARY KEY,
		    "inserted_at" TIMESTAMP DEFAULT current_timestamp,
		    "updated_at" TIMESTAMP DEFAULT current_timestamp
		  );
		`,
	},
	{
		1,
		"staff",
		`CREATE TABLE "staff" (
		    "id" BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
		    "email" VARCHAR(255) NOT NULL UNIQUE,
		    "password" VARCHAR NOT NULL,
		    "name" VARCHAR(25) NOT NULL UNIQUE,
		    "profile_image" VARCHAR,
		    "role" VARCHAR(255) DEFAULT 'unknown',
		    "inserted_at" TIMESTAMP DEFAULT current_timestamp,
		    "updated_at" TIMESTAMP DEFAULT current_timestamp,
		    CONSTRAINT fk_role FOREIGN KEY(role) REFERENCES staff_roles(name)
		  );
		`,
	},
	// {
	// 	1,
	// 	"staff_session",
	// 	`CREATE TABLE "staff_session" (
	// 	    "sid" VARCHAR NOT NULL COLLATE "default",
	// 	    "sess" JSON NOT NULL,
	// 	    "expire" timestamp(6) NOT NULL
	// 	  )
	// 	  WITH (OIDS=FALSE);
	// 	  ALTER TABLE "staff_session" ADD CONSTRAINT "staff_session_pkey" PRIMARY KEY ("sid") NOT DEFERRABLE INITIALLY IMMEDIATE;
	// 	  CREATE INDEX "IDX_staff_session_expire" ON "staff_session" ("expire");
	// 	`,
	// },
}
