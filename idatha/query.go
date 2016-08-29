package idatha

import (
	"net/http"

	"github.com/edgedagency/mulungu"
)

// Query used to generate executable query
type Query struct {
	Connection *mulungu.Connection
	Statement  map[string]string
}

// For used to add FOR condition to query
func (q *Query) For(subject string) {
	q.Statement["FOR"] = subject
}

// In adds IN clause to query
func (q *Query) In(subjects string) {
	q.Statement["IN"] = subjects
}

// Return adds return to query
func (q *Query) Return(statement string) {

}

// Bind variable to statement variables
func (q *Query) Bind(bindings map[string]string) {

}

// Filter adds filter conditions to query
func (q *Query) Filter(left, operator, right string) {

}

// Execute execute abitary statement on current connection
func (q *Query) Execute(database string, statement map[string]interface{}) (response map[string]interface{}, err error) {
	results, err := q.Connection.Execute(http.MethodPost, new(Dialect).Cursor(database), statement)
	if err != nil {
		return results, err
	}
	response = make(map[string]interface{})
	response["contents"] = results["result"]

	return response, nil
}
