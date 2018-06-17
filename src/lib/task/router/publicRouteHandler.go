package router

import (
	"encoding/json"
	"log"
	"net/http"

	"../schema"

	"github.com/softbrewery/gojoi/pkg/joi"

	"../service"
)

//TaskUsersAPIHandler  word api handler
func TaskUsersAPIHandler(w http.ResponseWriter, r *http.Request) {
	var words = service.UserList()
	writeStringArrayToResponse(w, words)
}

//RegisterUserHandler  word api handler
func RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	var response = service.RegisterUser(r.Body)
	writeStringArrayToResponse(w, response)
}

//UserInformationHandler User information retriever
func UserInformationHandler(w http.ResponseWriter, r *http.Request) {
	var email = r.URL.Query().Get("email")
	joi.Validate(email, schema.EmailQuery)
	var response = service.GetUserNameByEmail(email)
	writeStringArrayToResponse(w, response)
}

//AuthenticateUserHandler : Authenticates user
func AuthenticateUserHandler(w http.ResponseWriter, r *http.Request) {
	var email = r.FormValue("email")
	var password = r.FormValue("password")

	var response = service.AuthenticateUser(email, password)
	writeStringArrayToResponse(w, response)
}

//GetAllTasksForUserHandler : Authenticates user
func GetAllTasksForUserHandler(w http.ResponseWriter, r *http.Request) {
	var email = r.FormValue("email")

	service.GetAllTasksForUser(email)
	writeStringArrayToResponse(w, "response")
}

func writeStringArrayToResponse(w http.ResponseWriter, data interface{}) {
	jData, err := json.Marshal(data)
	if err != nil {
		log.Print("error", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jData)
}
