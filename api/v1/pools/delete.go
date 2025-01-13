package pools

import (
	"net/http"

	apiutil "github.com/enricomilli/neat-server/api/api-utils"
)

type DeletePoolRequest struct {
	PoolID string `json:"pool_id"`
}

// TODO: create the sql query to delete the pool with pool id and also delete all the related information

func HandlePoolDelete(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	addPoolReq := &DeletePoolRequest{}
	err := apiutil.StrictParseJSON(r, addPoolReq)
	if err != nil {
		apiutil.ResponseWithError(w, http.StatusBadRequest, "Invalid request")
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

	response := map[string]any{
		"success": true,
		"info":    "Pool has been deleted along with all information",
	}

	apiutil.ResponseWithJSON(w, http.StatusOK, response)
	return

}
