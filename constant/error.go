package constant

//ErrorCode represents an error code
type ErrorCode int

//Error codes constant
const (
	ErrDuplicate ErrorCode = iota - 1
	ErrFailedValidation
	ErrFailedBusinessRules
	ErrFailedEncryption
)
