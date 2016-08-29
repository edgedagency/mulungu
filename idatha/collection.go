/*
Package idatha mulungu provides a set of functions need to communicate with an arangodb database
*/
package idatha

import (
	"fmt"
	"net/http"

	"github.com/edgedagency/mulungu"
)

//Collection representation of a collection
type Collection struct {
	data map[string]interface{}
}

// Hydrate receives a map with string values which is used to hydrate KeyValue field of collection
func (c *Collection) Hydrate(data map[string]interface{}) {
	//initiate KeyValue store
	c.data = make(map[string]interface{})
	for key, value := range data {
		c.data[key] = value
	}
}

// Save collection object into database.
func (c *Collection) Save(data map[string]interface{}, database, collection string) error {
	fmt.Println("Collection save invoked")
	connection, _ := mulungu.NewConnection("http://127.0.0.1:8529", "root", "root")
	results, err := connection.Execute(http.MethodPost, new(Dialect).Create(database, collection), data)

	if err != nil {
		return err
	}

	fmt.Println(results)

	return nil
}

// Delete removes collection object from database.
func (c *Collection) Delete() error {
	return nil
}

// Update updates am existing collection object in database
func (c *Collection) Update(data map[string]interface{}, patch bool) error {
	if data != nil {
		c.Hydrate(data)
	}
	//depending on patch true or false/nil update or run patch function
	return nil
}

// NewQuery creates a new query object with current connection
func (c *Collection) NewQuery() *Query {
	connection, _ := mulungu.NewConnection("http://127.0.0.1:8529", "root", "root")
	return &Query{Connection: connection}
}
