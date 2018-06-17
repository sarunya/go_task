package router

import (
	"net/http"

	"../util"
)

//StartServer : Starts server and registers handler
func StartServer() {
	http.HandleFunc("/task", TaskUsersAPIHandler)
	http.HandleFunc("/registeruser", RegisterUserHandler)
	http.HandleFunc("/task/userinfo", UserInformationHandler)
	http.HandleFunc("/task/user/login", AuthenticateUserHandler)
	http.HandleFunc("/task/user/tasks", GetAllTasksForUserHandler)
	http.ListenAndServe("localhost:1952", util.BaseHTTPLog(http.DefaultServeMux))
}
