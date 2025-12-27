package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

type RequestBody struct {
	Index string `json:"index"`
}

func (h *handler) HandleDownload(w http.ResponseWriter, r *http.Request) {
	var body RequestBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		log.Printf("Error in getting Index")
	}
}
