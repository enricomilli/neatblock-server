package pools

import (
	"fmt"

	"github.com/enricomilli/neat-server/db"
)

// This method stores the current state of
// the pool struct into the db
func (pool *Pool) StorePoolStructState() error {

	database, err := db.NewClient()
	if err != nil {
		return fmt.Errorf("could not init db: %v", err)
	}

	query, values := db.BuildUpsertQuery("pools", pool, "id")

	_, err = database.NamedExec(query, values)
	if err != nil {
		return fmt.Errorf("could not store current pool state: %v", err)
	}

	return nil
}
