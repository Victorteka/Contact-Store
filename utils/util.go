package utils

import (
	"net/http"
	"encoding/json"
)

//Message func
func Message(status bool, message string)(map[string]interface{}){
	return map[string] interface{} {"status":status, "message":message}
}

//Respond func
func Respond(w http.ResponseWriter, data map[string] interface{}){
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}