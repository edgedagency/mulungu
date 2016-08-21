package mulungu

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Connection is a structure which will be used to represent a connection
type Connection struct {
	Host     string
	Client   *http.Client
	Username string
	Password string
}

// NewConnection creates connection and returns &pointer to refrence to this connection.
func NewConnection(host, username, password string) (connection *Connection, err error) {

	var results map[string]interface{}

	connection = &Connection{Host: host, Username: username, Password: password}
	connection.Client = new(http.Client)

	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", connection.Host, "/_api/version?details=true"), nil)

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(connection.Username, connection.Password)

	res, err := connection.Client.Do(req)

	if err != nil {
		panic("failed to establish connection")
	}

	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.New("failed to read content from body")
	}

	err = json.Unmarshal(b, &results)
	if err != nil {
		return nil, errors.New("failed to Unmarshal json results")
	}

	fmt.Println("---------------------")
	fmt.Println("raw results", results)
	fmt.Println("---------------------")

	errorNum, hasKey := results["errorNum"]
	if hasKey == true {
		errorString, _ := fmt.Printf("Failed to establish connection Host:%s ErrorNum:%g ErrorMsg:%s ", connection.Host, errorNum, results["errorMessage"])
		return nil, errors.New(string(errorString))
	}

	return connection, nil
}

// Execute will run abitary command on current connection.
func (c *Connection) Execute(statement map[string]interface{}) (results map[string]interface{}, err error) {

	compiledStatement, _ := json.Marshal(statement)
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s%s", c.Host, "/_api/cursor"), bytes.NewBuffer(compiledStatement))

	if err != nil {
		return nil, errors.New("Failed to create a request : " + err.Error())
	}

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(c.Username, c.Password)

	res, err := c.Client.Do(req)

	if err != nil {
		return nil, errors.New("Failed to run request : " + err.Error())
	}

	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.New("unable to read data from response : " + err.Error())
	}

	err = json.Unmarshal(b, &results)

	if err != nil {
		fmt.Println("Unmarshal of response body failed", err)
	}

	errorNum, hasKey := results["errorNum"]
	if hasKey == true {
		errorString, _ := fmt.Printf("Failed to execute query Query:%s Host:%s ErrorNum:%g ErrorMsg:%s ", compiledStatement, c.Host, errorNum, results["errorMessage"])
		return nil, errors.New(string(errorString))
	}

	return results, nil
}
