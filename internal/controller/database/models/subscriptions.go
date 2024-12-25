package models

var Subscriptions = []Schema{
	// {
	// 	0,
	// 	"subscriptions_type",
	// 	`CREATE TYPE subscription_type_recurrency AS ENUM
	// 	    ('month', 'one time');
	// 	  CREATE TABLE "subscriptions_type" (
	// 	    "id" BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	// 	    "name" VARCHAR,
	// 	    "amount" INT,
	// 	    "importance" INT DEFAULT 0,
	// 	    "active" BOOLEAN default FALSE,
	// 	    "recurrency" subscription_type_recurrency,
	// 	    "options" JSON DEFAULT ('[]'),
	// 	    "inserted_at" TIMESTAMP DEFAULT current_timestamp,
	// 	    "updated_at" TIMESTAMP DEFAULT current_timestamp
	// 	  );
	// 	`,
	// },
	{
		1,
		"subscriptions",
		`CREATE TYPE subscription_status AS ENUM
		    ('active', 'past_due', 'canceled', 'incomplete');
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
		  );
		`,
	},
}
