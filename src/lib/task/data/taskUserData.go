package data

import (
	"database/sql"
	"fmt"

	"../jsonstruct"
	"../jsonstruct/responsestruct"
	"../jsonstruct/tablestruct"
	"../util"

	sq "github.com/Masterminds/squirrel"
)

//mport "lib/dict/jsonstruct"me.logv2(componentName,'getProductsUrl',requestUrl);

const userTableName = "taskuser"
const email = "taskuser"

//GetAllUsersRegistered : Gets all words in pagination fashion
func GetAllUsersRegistered(limit int, offset int) []string {
	//query := jsonstruct.QueryOptions{"SELECT",["*"],userTableName};
	util.CreatePQClient()
	var query = fmt.Sprintf("SELECT %s from %s limit 20", email, userTableName)
	var word string
	fmt.Println("Execute query: ", query)
	dataRows, err := util.ExecuteQuery(query)
	util.ClosePQClient()
	var data = make([]string, 0)
	if err != nil {
		data[0] = "error"
	} else {
		fmt.Println("Execute query: done ")
		for dataRows.Next() {
			dataRows.Scan(&word)
			data = append(data, word)
		}
		fmt.Println(data, len(data), data[0])
	}
	return data
}

//RegisterTaskUser : Gets all words in pagination fashion
func RegisterTaskUser(hdata tablestruct.TaskUser) responsestruct.UserResp {
	//query := jsonstruct.QueryOptions{"SELECT",["*"],userTableName};
	util.CreatePQClient()

	record := tablestruct.TaskUserD{}
	record.Email = hdata.Email
	record.CreatedDate = hdata.CreatedDate
	record.ModifiedDate = hdata.ModifiedDate
	record.Data = string(hdata.Data)

	var query = fmt.Sprintf("INSERT INTO %s (email, data, created_date, modified_date) VALUES ('%s','%s','%s','%s')", userTableName, record.Email, record.Data, record.CreatedDate, record.ModifiedDate)

	fmt.Println("Execute query: ", query)
	dataRows, err := util.ExecuteQuery(query)
	return getUserResponse(dataRows, err)
}

//AuthenticateUser : Gets all words in pagination fashion
func AuthenticateUser(email string, password string) responsestruct.UserResp {
	queryOption := jsonstruct.QueryOptions{"SELECT", []string{"*"}, sq.Eq{}, userTableName}
	fmt.Println(util.ConstructQuery(queryOption))
	util.CreatePQClient()

	var query = fmt.Sprintf("SELECT data->>'name' from %s where email='%s' and password='%s'", userTableName, email, password)

	fmt.Println("Execute query: ", query)
	dataRows, err := util.ExecuteQuery(query)
	util.ClosePQClient()
	return getAuthenticateResponse(dataRows, err)
}

//GetUserDetails : Gets all words in pagination fashion
func GetUserDetails(email string, column string) string {
	//query := jsonstruct.QueryOptions{"SELECT",["*"],userTableName};
	util.CreatePQClient()

	var record string

	var query = fmt.Sprintf("SELECT data->>'%s' from %s where email='%s'", column, userTableName, email)

	fmt.Println("Execute query: ", query)
	dataRows, err := util.ExecuteQuery(query)
	util.ClosePQClient()
	if err != nil {
		fmt.Println("error it is", err.Error())
	} else {
		for dataRows.Next() {
			dataRows.Scan(&record)
			return record
		}
	}
	return "nil"
}

func getUserResponse(dataRow *sql.Rows, err error) responsestruct.UserResp {
	var response responsestruct.UserResp
	if err != nil {
		response = ConstructUserResponse(500, "", "InternalServerError", err.Error())
	} else {
		response = ConstructUserResponse(200, "RegistrationSuccess", "", "User Registered successfully")
	}
	return response
}

func getAuthenticateResponse(dataRow *sql.Rows, err error) responsestruct.UserResp {
	var response responsestruct.UserResp
	if err != nil || !dataRow.Next() {
		response = ConstructUserResponse(500, "", "InternalServerError", err.Error())
	} else {
		response = ConstructUserResponse(200, "AuthenticationSuccess", "", "User LoggedIn Successfully")
	}
	return response
}

//ConstructUserResponse : Constructs message as UserResp
func ConstructUserResponse(status int, message string, err string, desc string) responsestruct.UserResp {
	response := responsestruct.UserResp{}
	response.StatusCode = status
	response.Message = message
	response.Error = err
	response.Description = desc
	return response
}
