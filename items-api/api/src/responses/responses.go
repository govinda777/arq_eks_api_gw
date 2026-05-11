package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

// JSON returns a JSON response
func JSON(w http.ResponseWriter, statusCode int, dados interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(dados); err != nil {
		log.Fatal(err)
	}
}

func Err(w http.ResponseWriter, statusCode int, err error) {
	JSON(w, statusCode, struct {
		Err string `json:"err"`
	}{
		Err: err.Error(),
	})
}
