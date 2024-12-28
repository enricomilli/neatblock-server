package pools

import (
	"net/http"
	"time"

	apiutil "github.com/enricomilli/neat-server/api/api-utils"
	"github.com/google/uuid"
)

type AddPoolRequest struct {
	URL  string `json:"pool_url"`
	Name string `json:"pool_name"`
}

// TODO: add revenue share option
func HandleAddPool(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	addPoolReq := &AddPoolRequest{}
	err := apiutil.StrictParseJSON(r, addPoolReq)
	if err != nil {
		apiutil.ResponseWithError(w, http.StatusBadRequest, "Invalid request")
		return
	}

	// VALIDATING REQUEST
	pool_url := addPoolReq.URL
	if pool_url == "" {
		apiutil.ResponseWithError(w, http.StatusBadRequest, "No pool url found")
		return
	}

	poolName := addPoolReq.Name
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
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	newPool := Pool{
		ID:          uuid.NewString(),
		Status:      "pending",
		ObserverURL: pool_url,
		UserID:      userID,
		Name:        poolName,
		CreatedAt:   timestamp,
		UpdatedAt:   timestamp,
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

	err = newPool.StorePoolStructState()
	if err != nil {
		apiutil.ResponseWithError(w, http.StatusInternalServerError, err)
		return
	}

	err = newPool.ScrapeMiningData()
	if err != nil {
		apiutil.ResponseWithError(w, http.StatusInternalServerError, err)
		return
	}

	// make pool status as active after scraping rewards
	newPool.Status = "active"

	err = newPool.StorePoolStructState()
	if err != nil {
		apiutil.ResponseWithError(w, http.StatusInternalServerError, err)
		return
	}

	response := map[string]any{
		"success": true,
		"info":    "Pool was found and is now being tracked.",
	}

	apiutil.ResponseWithJSON(w, http.StatusOK, response)
	return
}
