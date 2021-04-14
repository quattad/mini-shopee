package responses

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// JSON is a function
// httpResponseWriter - interface used by HTTP handler to create HTTP response
// httpResponseWriter.WriteHeader(statusCode int) - sends HTTP response header with provided status code
// json.NewEncoder(w) - returns new encoder that writes to w
// type Encoder - writes JSON values to output stream
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)

	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

// ERROR is a function
func ERROR(w http.ResponseWriter, statusCode int, err error) {
	if err != nil {
		JSON(w, statusCode, struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		})

		return
	}

	JSON(w, http.StatusBadRequest, nil)
}
