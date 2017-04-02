package constant

//Status represents a status
type Status int

// status options
const (
	Active Status = iota + 1
	InActive
	Pending
	Deleted
	Warning
	Critical
)
