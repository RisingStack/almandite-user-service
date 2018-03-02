package repositories

import (
	"log"
	"os"
	"time"

	"github.com/go-pg/pg/orm"

	"github.com/RisingStack/almandite-user-service/models"

	"github.com/go-pg/pg"
)

var db *pg.DB
var dbLogger = log.New(os.Stdout, "PSQL ", log.Ldate|log.Ltime|log.LUTC)

// Users provide access ...
var Users UserRepository

// OpenConnection opens a connection to the DB
func OpenConnection(connURL string, debugSQL bool) error {
	pgOptions, err := pg.ParseURL(connURL)
	if err != nil {
		return err
	}

	db = pg.Connect(pgOptions)

	if debugSQL {
		db.OnQueryProcessed(func(event *pg.QueryProcessedEvent) {
			query, err := event.FormattedQuery()
			if err != nil {
				dbLogger.Println("Failed to format query", err)
			}

			dbLogger.Println(query, time.Since(event.StartTime))
		})
	}

	initRepositories()

	return nil
}

func initRepositories() {
	Users = NewUserRepository(db)
}

// Migrate does ...
func Migrate() error {
	// TODO: should this run in a transaction?
	for _, model := range []interface{}{&models.User{}} {
		err := db.CreateTable(model, &orm.CreateTableOptions{
			IfNotExists: true,
		})

		if err != nil {
			return err
		}
	}
	return nil
}

// CloseConnection closes the connection to the DB
func CloseConnection() {
	if err := db.Close(); err != nil {
		dbLogger.Println("Failed to close connection", err)
	} else {
		dbLogger.Println("Connection closed")
	}
}
