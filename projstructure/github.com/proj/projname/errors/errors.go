package errors

import "fmt"

var RecordNotFound = fmt.Errorf("Record not found")



type TaskStatusType string

const (
	Created   TaskStatusType = "Created"
	Pending   TaskStatusType = "Pending"
	SpillOver TaskStatusType = "SpillOver"
	Completed TaskStatusType = "Completed"
)
