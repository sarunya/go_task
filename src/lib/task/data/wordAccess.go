package data

import (
	"encoding/json"
	"fmt"
	"../jsonstruct"
	"../util"
)

//mport "lib/dict/jsonstruct"me.logv2(componentName,'getProductsUrl',requestUrl);

//tableName = "wn_synset"

//GetAllWordsWithPagination : Gets all words in pagination fashion
func GetAllWordsWithPagination(limit int, offset int) []string {
	//query := jsonstruct.QueryOptions{"SELECT",["*"],tableName};
	util.CreatePQClient()
	var query = "SELECT word from wn_synset limit 20"
	var word string
	fmt.Println("Execute query: ", query)
	dataRows := util.ExecuteQuery(query)
	fmt.Println("Execute query: done ")
	var data = make([]string, 0)
	for dataRows.Next() {
		dataRows.Scan(&word)
		data = append(data, word)
	}
	util.ClosePQClient()
	fmt.Println(data, len(data), data[0])
	var jsonData = getDataFromWordsTable(data[0])
	fmt.Println(jsonData)
	return data
}

//GetMeaningForWord : Returns meaning for the given word
func GetMeaningForWord(word string) string {
	util.CreatePQClient()
	var query = "SELECT word from wn_gloss limit 20"
	fmt.Println("Execute query: ", query)
	dataRows := util.ExecuteQuery(query)
	fmt.Println("Execute query: done ")
	var data = make([]string, 20)
	for dataRows.Next() {
		dataRows.Scan(&word)
		data = append(data, word)
	}
	util.ClosePQClient()
	return "data"
}

func getDataFromWordsTable(getword string) jsonstruct.WordMeaning {
	//query := jsonstruct.QueryOptions{"SELECT",["*"],tableName};
	util.CreatePQClient()
	var query = "SELECT synset_id from wn_synset where word='%s';"
	query = fmt.Sprintf(query, getword)
	fmt.Println(query)
	dataRows := util.ExecuteQuery(query)
	fmt.Println("Execute query: done ")
	var data jsonstruct.WordMeaning
	var similarWords = make([]string, 0)
	var dataRow interface{}
	var str string
	if dataRows.Next() {
		data = jsonstruct.WordMeaning{
			Word:         getword,
			WordID:       ",",
			Meaning:      "string",
			SimilarWords: similarWords}
		dataRows.Scan(&dataRow)
	}

	var dataR []byte
	json.Unmarshal(dataR, dataRow)
	fmt.Println(str, "data row", dataR)
	util.ClosePQClient()
	return data
}
