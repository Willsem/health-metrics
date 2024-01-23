-- Create "users" table
CREATE TABLE "users" ("uuid" uuid NOT NULL, "email" character varying NOT NULL, "login" character varying NOT NULL, "password" character varying NOT NULL, PRIMARY KEY ("uuid"));
-- Create index "users_email_key" to table: "users"
CREATE UNIQUE INDEX "users_email_key" ON "users" ("email");
-- Create index "users_login_key" to table: "users"
CREATE UNIQUE INDEX "users_login_key" ON "users" ("login");
-- Create index "user_email" to table: "users"
CREATE UNIQUE INDEX "user_email" ON "users" ("email");
-- Create index "user_login" to table: "users"
CREATE UNIQUE INDEX "user_login" ON "users" ("login");
-- Create "accesses" table
CREATE TABLE "accesses" ("uuid" uuid NOT NULL, "user_given_accesses" uuid NULL, "user_recieved_accesses" uuid NULL, PRIMARY KEY ("uuid"), CONSTRAINT "accesses_users_given_accesses" FOREIGN KEY ("user_given_accesses") REFERENCES "users" ("uuid") ON DELETE SET NULL, CONSTRAINT "accesses_users_recieved_accesses" FOREIGN KEY ("user_recieved_accesses") REFERENCES "users" ("uuid") ON DELETE SET NULL);
-- Create "metrics" table
CREATE TABLE "metrics" ("uuid" uuid NOT NULL, "metric_type" character varying NOT NULL, "value" bigint NOT NULL, "date" timestamptz NOT NULL, "user_metrics" uuid NULL, PRIMARY KEY ("uuid"), CONSTRAINT "metrics_users_metrics" FOREIGN KEY ("user_metrics") REFERENCES "users" ("uuid") ON DELETE SET NULL);
-- Create index "metric_metric_type_user_metrics" to table: "metrics"
CREATE INDEX "metric_metric_type_user_metrics" ON "metrics" ("metric_type", "user_metrics");
