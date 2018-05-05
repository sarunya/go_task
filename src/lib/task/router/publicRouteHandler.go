package router

import (
	"encoding/json"
	"log"
	"net/http"

	"../service"
)

//WordAPIHandler  word api handler
func WordAPIHandler(w http.ResponseWriter, r *http.Request) {
	var words = service.WordsList()
	writeStringArrayToResponse(w, words)
}

func writeStringArrayToResponse(w http.ResponseWriter, data interface{}) {
	jData, err := json.Marshal(data)
	if err != nil {
		log.Print("error", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jData)
}
