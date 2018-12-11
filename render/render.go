package render

import (
	"encoding/json"
	"net/http"
)

// JSON sets a json http header, encodes data to json and write to response writer.
func JSON(w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)

	err := encoder.Encode(data)
	if err != nil {
		http.Error(w, `{"error": "unable to encode json."}`, http.StatusInternalServerError)
		return err
	}

	return nil
}
