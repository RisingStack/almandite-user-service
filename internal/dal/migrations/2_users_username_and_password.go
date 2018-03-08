package migrations

import (
	"github.com/go-pg/migrations"
)

func init() {
	migrations.Register(func(db migrations.DB) error {
		_, err := db.Exec(`
			ALTER TABLE IF EXISTS "users"
				ADD COLUMN "username" text NOT NULL,
				ADD COLUMN "password" text NOT NULL,
				ADD CONSTRAINT username_unique UNIQUE ("username");
			`)

		return err
	}, func(db migrations.DB) error {
		_, err := db.Exec(`
			ALTER TABLE IF EXISTS "users"
				DROP COLUMN IF EXISTS "username",
				DROP COLUMN IF EXISTS "password",
 				DROP CONSTRAINT IF EXISTS username_unique;
			`)

		return err
	})
}
