package service

import (
	"encoding/json"
	"io"
	"io/ioutil"

	"../data"
	"../jsonstruct/responsestruct"
	"../jsonstruct/tablestruct"
)

//UserList : return words list
func UserList() []string {
	return data.GetAllUsersRegistered(10, 20)
}

//RegisterUser : return words list
func RegisterUser(body io.Reader) responsestruct.UserResp {
	hdata := tablestruct.TaskUser{}
	readbody, _ := ioutil.ReadAll(body)
	json.Unmarshal(readbody, &hdata)
	return data.RegisterTaskUser(hdata)
}

//AuthenticateUser : return words list
func AuthenticateUser(email string, password string) responsestruct.UserResp {
	name := GetUserNameByEmail(email)

	if name == "nil" || name == "" {
		return data.ConstructUserResponse(404, "", "UserNotFound", "User is not found")
	} else {
		return data.AuthenticateUser(email, password)
	}
}

//GetUserNameByEmail : returns user information
func GetUserNameByEmail(email string) string {
	return data.GetUserDetails(email, "name")
}
