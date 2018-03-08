package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"

	"github.com/RisingStack/almandite-user-service/internal/config"
	"github.com/RisingStack/almandite-user-service/internal/dal"
	"github.com/go-pg/migrations"
	"github.com/go-pg/pg"

	_ "github.com/RisingStack/almandite-user-service/internal/dal/migrations"
)

const usageText = `This program runs command on the db. Supported commands are:
  - up - runs all available migrations.
  - down - reverts last migration.
  - reset - reverts all migrations.
  - version - prints current db version.
  - set_version [version] - sets db version without running migrations.
`

func main() {
	flag.Usage = usage
	flag.Parse()

	fmt.Println(flag.Args())

	db := dal.NewDAL()

	configuration := config.GetConfiguration()

	if err := db.OpenConnection(
		configuration.PostgresURL,
		configuration.DebugSQL,
	); err != nil {
		log.Fatal("Failed to open DB conenction", err)
	}

	err := db.DB().RunInTransaction(func(tx *pg.Tx) error {
		oldV, newV, err := migrations.Run(tx, flag.Args()...)
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

func usage() {
	fmt.Printf(usageText)
	flag.PrintDefaults()
	os.Exit(2)
}
