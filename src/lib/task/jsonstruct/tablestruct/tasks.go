package tablestruct

import "encoding/json"

//Task : represents synset tagble
type Task struct {
	TaskTile      string
	ScheduledDate string
	Tag           string
	IsCompleted   string
	Comments      string
}

//TaskData : represents synset tagble
type TaskData struct {
	ID           string
	Data         json.RawMessage
	CreatedDate  string
	ModifiedDate string
}

//TaskDataD : represents synset tagble
type TaskDataD struct {
	ID           string
	Data         string
	CreatedDate  string
	ModifiedDate string
}
