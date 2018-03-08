package migrations

import (
	"github.com/go-pg/migrations"
)

func init() {
	migrations.Register(func(db migrations.DB) error {
		_, err := db.Exec(`
			CREATE TABLE IF NOT EXISTS "users" ("id" bigserial, "first_name" text, "last_name" text, PRIMARY KEY ("id"));
			CREATE TABLE IF NOT EXISTS "access_logs" (
				"id" bigserial, "user_id" bigint REFERENCES "users" ("id"), "timestamp" timestamp, "ip_address" inet, "event" text,
				PRIMARY KEY ("id")
				)
			`)

		return err
	}, func(db migrations.DB) error {
		_, err := db.Exec(`
			DROP TABLE IF EXISTS "access_logs";
			DROP TABLE IF EXISTS "users"
			`)

		return err
	})
}
