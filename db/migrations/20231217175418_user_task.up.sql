CREATE TABLE "users" (
  "username" varchar PRIMARY KEY,
  "full_name" varchar NOT NULL,
  "hashed_password" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "password_changed_at" timestamptz DEFAULT '0001-01-01 00:00:00.00Z',
  "created_at" timestamptz DEFAULT 'now()'
);

CREATE TABLE "tasks" (
  "id" bigserial PRIMARY KEY,
  "username" varchar NOT NULL,
  "title" varchar,
  "description" varchar,
  "due_date" timestamptz,
  "priority" int,
  "created_at" timestamptz DEFAULT 'now()'
);

CREATE INDEX ON "users" ("username");

CREATE INDEX ON "users" ("email");

ALTER TABLE "tasks" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");
