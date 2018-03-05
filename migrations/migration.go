package migrations

import (
	"fmt"

	"github.com/go-pg/migrations"
	"github.com/go-pg/pg"
)

// Run ...
func Run(tx *pg.Tx) error {
	oldV, newV, err := migrations.Run(tx, "up")
	if err != nil {
		return err
	}

	fmt.Printf("Migrations ran. Updated from %v to %v\n", oldV, newV)
	return nil
}
