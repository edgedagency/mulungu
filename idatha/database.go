package idatha

import (
	"net/http"

	"github.com/edgedagency/mulungu"
)

// Database struct used to access database operations
type Database struct {
	Connection *mulungu.Connection
}

//Create create a new database
func (d *Database) Create(name string) (results map[string]interface{}, err error) {
	return d.Connection.Execute(http.MethodPost, new(Dialect).CreateDatabase(), map[string]interface{}{name: name})
}
