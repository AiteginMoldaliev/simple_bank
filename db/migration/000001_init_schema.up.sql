CREATE TABLE "accounts" (
	"id" bigserial PRIMARY KEY,
	"owner" VARCHAR not null,
	"balance" bigint not null,
	"currency" varchar not null,
	"created_at" TIMESTAMPTZ not null DEFAULT (now())
);

CREATE TABLE "entries" (
	"id" bigserial PRIMARY KEY, 
	"account_id" BIGINT not null,
	"amount" BIGINT not NULL,
	"created_at" TIMESTAMPTZ not null DEFAULT (now())
);

CREATE TABLE "transfers" (
	"id" bigserial PRIMARY KEY,
	"from_account_id" BIGINT not null, 
	"to_account_id" BIGINT not null,
	"amount" BIGINT not NULL,
	"created_at" TIMESTAMPTZ not null DEFAULT (now())
);

ALTER TABLE "entries" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("from_account_id") REFERENCES "accounts" ("id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("to_account_id") REFERENCES "accounts" ("id");

CREATE INDEX ON "accounts" ("owner");

CREATE INDEX ON "entries" ("account_id");

CREATE INDEX ON "transfers" ("from_account_id");

CREATE INDEX ON "transfers" ("to_account_id");

CREATE INDEX ON "transfers" ("from_account_id", "to_account_id");

COMMENT ON COLUMN "entries"."amount" IS 'can be negative or positive';
 
COMMENT on COLUMN "transfers"."amount" IS 'must be positive';