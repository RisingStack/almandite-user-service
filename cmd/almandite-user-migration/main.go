package main

import (
	"fmt"
	"log"

	_ "github.com/joho/godotenv/autoload"

	"github.com/RisingStack/almandite-user-service/config"
	"github.com/RisingStack/almandite-user-service/dal"
	"github.com/go-pg/migrations"
	"github.com/go-pg/pg"

	_ "github.com/RisingStack/almandite-user-service/migrations"
)

func main() {
	db := dal.NewDAL()

	configuration := config.GetConfiguration()

	if err := db.OpenConnection(
		configuration.PostgresURL,
		configuration.DebugSQL,
	); err != nil {
		log.Fatal("Failed to open DB conenction", err)
	}

	err := db.DB().RunInTransaction(func(tx *pg.Tx) error {
		oldV, newV, err := migrations.Run(tx, "up")
		if err != nil {
			return err
		}

		fmt.Printf("Migrations ran. Updated from %v to %v\n", oldV, newV)
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
}
