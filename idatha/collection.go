/*
Package idatha mulungu provides a set of functions need to communicate with an arangodb database
*/
package idatha

import "edged.agency/mulungu"

//Connection currently established connection
var connection *mulungu.Connection
var err error

func init() {
	connection, err = mulungu.NewConnection("http://127.0.0.1:8529", "root", "root")
	if err != nil {
		panic(err)
	}
}

// Collection is a respresentation of an arangodb collection
// will allow interaction with arangodb to perform CRUD and enable execution of AQL commands.
type Collection struct {
	ID  string `json:"_id"`
	Key string `json:"_key"`
	Rev string `json:"_rev"`
	//query Query
	Data interface{}
}

func (c *Collection) hydrate(data map[string]interface{}) {

}

// Save collection object into database.
func (c *Collection) Save(data map[string]interface{}) error {

	return nil
}

// Delete removes collection object from database.
func (c *Collection) Delete() error {
	return nil
}

// Update updates am existing collection object in database
func (c *Collection) Update(data map[string]interface{}, override bool) error {
	return nil
}

// NewQuery creates a new query object with current connection
func (c *Collection) NewQuery() *Query {
	return &Query{Connection: connection}
}
