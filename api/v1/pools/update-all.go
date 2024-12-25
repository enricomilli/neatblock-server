package pools

import (
	"fmt"
	"net/http"

	apiutil "github.com/enricomilli/neat-server/api/api-utils"
	"github.com/enricomilli/neat-server/db"
)

// Scrapes the mining data for all pools in the database
// TODO: make it a longer running process where the pool requests are made at random intervals
func HandleUpdateAll(w http.ResponseWriter, r *http.Request) {

	database, err := db.NewClient()
	if err != nil {
		apiutil.ResponseWithError(w, http.StatusInternalServerError, err)
		return
	}

	query := `
		select * from pools
	`

	allPools := []Pool{}
	err = database.Select(&allPools, query)
	if err != nil {
		apiutil.ResponseWithError(w, http.StatusInternalServerError, "could not get all pools:", err)
		return
	}

	// I think this will create a race condition is i make it multi threaded
	for _, pool := range allPools {
		err := pool.ScrapeMiningData()
		if err != nil {
			fmt.Printf("Could not scrape pool: %s\nError: %v", pool.Name, err)
		}
	}

	return
}
