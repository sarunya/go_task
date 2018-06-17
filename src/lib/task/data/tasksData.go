package data

import (
	"fmt"

	"../jsonstruct/responsestruct"
	"../jsonstruct/tablestruct"
	"../util"
)

const tasksTableName = "taskuser"

var tasks = [...]string{"", "saru"}

//GetAllTasksForUser : Gets all words in pagination fashion
func GetAllTasksForUser(limit int, offset int) []string {
	//query := jsonstruct.QueryOptions{"SELECT",["*"],userTableName};
	util.CreatePQClient()
	var query = fmt.Sprintf("SELECT data from %s where data->>'email'='%s'", tasksTableName, email)
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

//AddTaskForUser : Gets all words in pagination fashion
func AddTaskForUser(hdata tablestruct.TaskData) responsestruct.UserResp {
	//query := jsonstruct.QueryOptions{"SELECT",["*"],userTableName};
	util.CreatePQClient()

	record := tablestruct.TaskDataD{}
	record.ID = hdata.ID
	record.CreatedDate = hdata.CreatedDate
	record.ModifiedDate = hdata.ModifiedDate
	record.Data = string(hdata.Data)

	var query = fmt.Sprintf("INSERT INTO %s (id, data, created_date, modified_date) VALUES ('%s','%s','%s','%s')", userTableName, record.ID, record.Data, record.CreatedDate, record.ModifiedDate)

	fmt.Println("Execute query: ", query)
	dataRows, err := util.ExecuteQuery(query)
	return getUserResponse(dataRows, err)
}
