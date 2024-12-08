package apiutil

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/enricomilli/neat-server/msg"
)

func StrictParseJSON(r *http.Request, v any) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields() // Ensures strict parsing
	if err := decoder.Decode(v); err != nil {
		return err
	}
	return nil
}

// ResponseWithJSON sends a JSON response with the given status code and payload
func ResponseWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	fmt.Printf("JSON Response: %+v\n", payload)

	response, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling JSON response: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_, err = w.Write(response)
	if err != nil {
		log.Printf("Error writing response: %v", err)
	}
}

// ErrorResponse represents the structure of an error response
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message,omitempty"`
	Code    int    `json:"code"`
}

// ResponseWithError sends a JSON error response with the given status code and error message
func ResponseWithError(w http.ResponseWriter, statusCode int, args ...interface{}) {
	var parts []string

	for _, arg := range args {
		switch v := arg.(type) {
		case string:
			parts = append(parts, v)
		case error:
			parts = append(parts, v.Error())
		}
	}

	message := ""
	for i, part := range parts {
		if i > 0 && !strings.HasSuffix(parts[i-1], " ") && !strings.HasPrefix(part, " ") {
			message += " "
		}
		message += part
	}

	msg.MsgTelegram("Error from Neatblock API: \n" + message)
	errorResponse := ErrorResponse{
		Error:   http.StatusText(statusCode),
		Message: message,
		Code:    statusCode,
	}
	ResponseWithJSON(w, statusCode, errorResponse)
}
