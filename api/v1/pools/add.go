package pools

import (
	"net/http"
	"os"

	apiutil "github.com/enricomilli/neat-server/api/api-utils"
	"github.com/enricomilli/neat-server/db"
	"github.com/supabase-community/supabase-go"
)

func HandleAddPool(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	// VALIDATING REQUEST
	pool_url := r.URL.Query().Get("pool_url")
	if pool_url == "" {
		apiutil.ResponseWithError(w, http.StatusBadRequest, "No pool url found")
		return
	}

	poolName := r.URL.Query().Get("name")
	if poolName == "" {
		apiutil.ResponseWithError(w, http.StatusBadRequest, "No pool name found")
		return
	}

	if apiutil.IsValidURL(pool_url) == false {
		apiutil.ResponseWithError(w, http.StatusBadRequest, "Not a valid url")
		return
	}

	userID := ctx.Value("userID").(string)
	if userID == "" {
		apiutil.ResponseWithError(w, http.StatusBadRequest, "No user found in ctx")
		return
	}

	userToken := ctx.Value("token").(string)
	if userToken == "" {
		apiutil.ResponseWithError(w, http.StatusUnauthorized, "No auth token found, signout and sign back in.")
		return
	}

	// VALIDATING WITH POOL PROVIDER
	newPool := Pool{
		ObserverURL: pool_url,
		OwnerID:     userID,
		Name:        poolName,
	}

	provider, err := newPool.NewProviderInterface()
	if err != nil {
		apiutil.ResponseWithError(w, http.StatusBadRequest, err)
		return
	}

	err = provider.ValidateURL(pool_url)
	if err != nil {
		apiutil.ResponseWithError(w, http.StatusBadRequest, err)
		return
	}

	err = newPool.UpdatePoolData()
	if err != nil {
		apiutil.ResponseWithError(w, http.StatusInternalServerError, err)
		return
	}

	// STORING THE NEW POOL TO PERSISTANT STORAGE
	sbClient, err := supabase.NewClient(os.Getenv("SUPABASE_URL"), userToken, &supabase.ClientOptions{})
	if err != nil {
		apiutil.ResponseWithError(w, http.StatusInternalServerError, err)
		return
	}

	_, err = sbClient.From("pools").Upsert(newPool, "pool_url", "*", "exact").ExecuteTo(&newPool)
	if err != nil {
		code, msg := db.HandleSupabaseError(err)
		apiutil.ResponseWithError(w, code, msg)
		return
	}

	response := map[string]any{
		"success": true,
		"info":    "Pool was found and is now being tracked.",
	}

	apiutil.ResponseWithJSON(w, http.StatusOK, response)
	return
}
