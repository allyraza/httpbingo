package httpbin

import (
	"encoding/json"
	"log"
	"net/http"
)

func (h *Httpbin) JSON(w http.ResponseWriter, data interface{}) {
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Printf("render json: %v\n", err)
	}
}
