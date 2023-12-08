ALTER TABLE IF EXISTS "accounts" DROP constraint if EXISTS "owner_currency_key";

ALTER TABLE IF EXISTS "accounts" DROP constraint if EXISTS "owner_currency_fkey";

DROP TABLE IF EXISTS "users";