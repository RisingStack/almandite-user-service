package dal

import (
	"log"
	"os"
	"time"

	"github.com/RisingStack/almandite-user-service/migrations"

	"github.com/go-pg/pg"
)

// DAL (Data Access Layer)
type DAL interface {
	OpenConnection(connURL string, debugSQL bool) error
	Migrate() error
	CloseConnection()
	Users() UserRepository
	AccessLogs() AccessLogRepository
}

type dal struct {
	db         *pg.DB
	logger     *log.Logger
	users      UserRepository
	accessLogs AccessLogRepository
}

// NewDAL returns an implementation for the DAL interface
func NewDAL() DAL {
	return &dal{
		logger: log.New(os.Stdout, "PSQL ", log.Ldate|log.Ltime|log.LUTC),
	}
}

// OpenConnection opens a connection to the DB
func (d *dal) OpenConnection(connURL string, debugSQL bool) error {
	pgOptions, err := pg.ParseURL(connURL)
	if err != nil {
		return err
	}

	d.db = pg.Connect(pgOptions)

	if debugSQL {
		d.db.OnQueryProcessed(func(event *pg.QueryProcessedEvent) {
			query, err := event.FormattedQuery()
			if err != nil {
				d.logger.Println("Failed to format query", err)
			}

			d.logger.Println(query, time.Since(event.StartTime))
		})
	}

	d.users = newUserRepository(d.db)
	d.accessLogs = newAccessLogRepository(d.db)

	return nil
}

// Migrate does ...
func (d *dal) Migrate() error {
	err := d.db.RunInTransaction(func(tx *pg.Tx) error {
		return migrations.Run(tx)
	})

	if err != nil {
		return err
	}

	return nil
}

// CloseConnection closes the connection to the DB
func (d *dal) CloseConnection() {
	if err := d.db.Close(); err != nil {
		d.logger.Println("Failed to close connection", err)
	} else {
		d.logger.Println("Connection closed")
	}
}

func (d *dal) Users() UserRepository {
	return d.users
}

func (d *dal) AccessLogs() AccessLogRepository {
	return d.accessLogs
}
