package model

//Status represents a status
type Status int

// status options
const (
	Active Status = iota
	InActive
	Pending
	Deleted
	Warning
	Critical
)
