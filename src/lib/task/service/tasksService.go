package service

import "../data"

//GetAllTasksForUser : Gets all taks for given user email
func GetAllTasksForUser(email string) {
	data.GetAllTasksForUser(email)
}
