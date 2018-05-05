package service

import (
	"../data"
)

//WordsList : return words list
func WordsList() []string {
	return data.GetAllWordsWithPagination(10, 20)
}
