package util

import (
	"database/sql"
	"fmt"
	"log"

	//comment
	_ "github.com/lib/pq"
)

var connStr = "user=postgres dbname=wn_pro_mysql sslmode=disable"
var dbConn *sql.DB

//CreatePQClient : Creates PQ client
func CreatePQClient() {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	if db != nil {
		log.Println("DB Connection is established!")
	}
	dbConn = db
}

//ExecuteQuery : Execute query
func ExecuteQuery(query string) *sql.Rows {
	fmt.Print("query: ", query)
	rows, err := dbConn.Query(query)
	if err != nil {
		log.Printf("error?? yes it is")
		log.Fatal(err)
	}
	return rows
}

//ClosePQClient : Close client
func ClosePQClient() {
	err := dbConn.Close()
	if err != nil {
		log.Fatal(err)
	}
}

// func constructQuery(queryOption json) {

// }
