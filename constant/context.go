package constant

//ContextKey constant to be used when setting context values
type ContextKey int

//Context constants starting at 1000
const (
	ContextUser ContextKey = iota + 1000
)
