package enums

import "strings"

type (
	Status string
)

const (
	// Only 2 status messages are valid
	InProgress Status = "IN_PROGRESS"
	Completed  Status = "COMPLETED"
)

func (s Status) IsValid() bool {
	switch s.ToUpper() {
	case InProgress, Completed:
		return true
	}
	return false
}

func (s Status) ToUpper() Status {
	return Status(strings.ToUpper(string(s)))
}
