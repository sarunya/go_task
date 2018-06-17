package tablestruct

import (
	"encoding/json"
)

//TaskUser : represents synset tagble
type TaskUser struct {
	Email        string
	Data         json.RawMessage
	CreatedDate  string
	ModifiedDate string
}

//TaskUserD : represents synset tagble
type TaskUserD struct {
	Email        string
	Data         string
	CreatedDate  string
	ModifiedDate string
}
