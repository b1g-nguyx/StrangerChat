CREATE TABLE "reports" (
  "id" varchar PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "reporter_id" varchar NOT NULL,
  "reported_id" varchar NOT NULL,
  "reason" text NOT NULL,
  "status" varchar NOT NULL DEFAULT 'PENDING',
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "deleted_at" timestamptz
);

CREATE TABLE "report_evidences" (
  "id" varchar PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "report_id" varchar NOT NULL,
  "room_id" varchar NOT NULL,
  "chat_logs" jsonb NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "deleted_at" timestamptz
);

ALTER TABLE "reports" ADD FOREIGN KEY ("reporter_id") REFERENCES "users" ("id");
ALTER TABLE "reports" ADD FOREIGN KEY ("reported_id") REFERENCES "users" ("id");
ALTER TABLE "report_evidences" ADD FOREIGN KEY ("report_id") REFERENCES "reports" ("id");
ALTER TABLE "report_evidences" ADD FOREIGN KEY ("room_id") REFERENCES "rooms" ("id");
