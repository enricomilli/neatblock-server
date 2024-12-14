package v2pools

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	apiutil "github.com/enricomilli/neat-server/api/api-utils"
)

func ForemanTests(w http.ResponseWriter, r *http.Request) {

	client := &http.Client{}
	// clientId := os.Getenv("FOREMAN_CLIENT_ID")
	baseURL := "https://api.foreman.mn/api/v2"
	// endpoint := "/clients/" + testClientId + "/group]"
	endpoint := "/pool-keys"
	url := baseURL + endpoint

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		apiutil.ResponseWithError(w, http.StatusInternalServerError, "could not create request:", err)
		return
	}
	req.Header.Add("Authorization", "Token "+os.Getenv("FOREMAN_API_KEY"))
	req.Header.Add("Accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		apiutil.ResponseWithError(w, http.StatusBadRequest, "could not complete request:", err)
		return
	}
	defer res.Body.Close()

	contentType := res.Header.Get("Content-Type")
	fmt.Println("Content Type:", contentType)

	if strings.Contains(contentType, "application/json") {
		var body interface{}
		err = json.NewDecoder(res.Body).Decode(&body)
		if err != nil {
			apiutil.ResponseWithError(w, http.StatusInternalServerError, "could not parse json res:", err)
			return
		}

		fmt.Printf("%+v\n", body)
		fmt.Fprintf(w, "%+v\n", body)
		return
	} else if strings.Contains(contentType, "text/html") {

		// Read HTML response body
		htmlBytes, err := io.ReadAll(res.Body)
		if err != nil {
			apiutil.ResponseWithError(w, http.StatusInternalServerError, "could not read html response:", err)
			return
		}

		// Convert bytes to string
		htmlContent := string(htmlBytes)

		// You can either:
		// 1. Forward the HTML content directly
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprint(w, htmlContent)

		return
	}

}
