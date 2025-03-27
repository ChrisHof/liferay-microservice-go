package main

import (
	"encoding/json"
	"net/http"
)

func InitHttpHandlers() {
	http.HandleFunc("/liveness", Liveness)
	http.HandleFunc("/object-action", ObjectAction)
}

func Liveness(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}

func ObjectAction(w http.ResponseWriter, r *http.Request) {
	if IsValidRequest(w, r) {
		var jsonBody map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&jsonBody)
		if err != nil {
			Logger.Println("Invalid JSON:", r.Body)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		Logger.Println(jsonBody)
		w.WriteHeader(http.StatusOK)
	}
}
