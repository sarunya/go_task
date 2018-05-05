package router

import (
	"../util"
	"net/http"
)

//StartServer : Starts server and registers handler
func StartServer() {
	http.HandleFunc("/word", WordAPIHandler)
	http.ListenAndServe("localhost:1952", util.BaseHTTPLog(http.DefaultServeMux))
}
