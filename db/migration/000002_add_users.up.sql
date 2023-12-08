CREATE TABLE "users" (
	"username" varchar PRIMARY KEY,
	"hashed_password" varchar NOT NULL,
	"full_name"	varchar NOT NULL,
	"email"	varchar UNIQUE NOT NULL,
	"password_changed_at" TIMESTAMPTZ not null DEFAULT '0001-01-01 00:00:00Z',
	"created_at" TIMESTAMPTZ not null DEFAULT (now())
);

ALTER TABLE "accounts" ADD FOREIGN KEY ("owner") REFERENCES "users"	("username");

-- CREATE UNIQUE INDEX ON "accounts" ("owner", "currency"); 

ALTER TABLE "accounts" ADD constraint "owner_currency_key" UNIQUE ("owner", "currency");