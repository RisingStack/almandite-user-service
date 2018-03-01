package main

import (
	"log"
	"os"
	"time"

	"github.com/go-pg/pg/orm"

	"github.com/RisingStack/almandite-user-service/models"
	"github.com/RisingStack/almandite-user-service/repositories"

	"github.com/go-pg/pg"
)

var db *pg.DB
var dbLogger = log.New(os.Stdout, "PSQL ", log.Ldate|log.Ltime|log.LUTC)

// UserRepository ...
var UserRepository repositories.UserRepository

// OpenDBConnection opens a connection to the DB
func OpenDBConnection() {
	// TODO: config or environment variables
	db = pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "w7o4bvt8ncp0ksd",
		Database: "almandite",
	})

	// TODO: make query log configurable
	db.OnQueryProcessed(func(event *pg.QueryProcessedEvent) {
		query, err := event.FormattedQuery()
		if err != nil {
			dbLogger.Println("Failed to format query", err)
		}

		dbLogger.Println(query, time.Since(event.StartTime))
	})

	initRepositories()
}

func initRepositories() {
	UserRepository = repositories.NewUserRepository(db)
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

// CloseDBConnection closes the connection to the DB
func CloseDBConnection() {
	if err := db.Close(); err != nil {
		dbLogger.Println("Failed to close connection", err)
	} else {
		dbLogger.Println("Connection closed")
	}
}
