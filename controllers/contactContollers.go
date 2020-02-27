package controllers

import (
	"encoding/json"
	"fmt"
	"go-contacts/models"
	u "go-contacts/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//CreateContact controller
var CreateContact = func(w http.ResponseWriter, r *http.Request) {
	//Grab the id of the user that send the request
	user := r.Context().Value("user").(uint)
	contact := &models.Contact{}

	err := json.NewDecoder(r.Body).Decode(contact)

	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}
	contact.UserID = user
	resp := contact.Create()
	u.Respond(w, resp)
}

//GetContactForUser for user
var GetContactForUser = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		u.Respond(w, u.Message(false, "There was an error in your request"))
		fmt.Println(err)
		return
	}

	data := models.GetContacts(uint(id))

	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)

}
