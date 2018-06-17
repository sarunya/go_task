package util

import (
	"database/sql"
	"fmt"
	"log"

	//comment

	"../jsonstruct"
	_ "github.com/lib/pq"
)

var connStr = "user=postgres dbname=sar_tasks sslmode=disable"
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
	//ConstructQuery()
}

//ExecuteQuery : Execute query
func ExecuteQuery(query string) (*sql.Rows, error) {
	fmt.Print("query: ", query)
	rows, err := dbConn.Query(query)
	return rows, err
}

//ExecuteQueryBuilder : Executes the query builder
func ExecuteQueryBuilder(queryOption jsonstruct.QueryOptions) (*sql.Rows, error) {
	CreatePQClient()
	queryBuilder := ConstructQueryBuilder(queryOption)
	fmt.Println(queryBuilder.ToSql())
	rows, err := queryBuilder.RunWith(dbConn).Query()
	ClosePQClient()
	return rows, err
}

//ClosePQClient : Close client
func ClosePQClient() {
	err := dbConn.Close()
	if err != nil {
		log.Fatal(err)
	}
}
