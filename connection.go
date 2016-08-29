package mulungu

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/edgedagency/mulungu/util"
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
	connection = &Connection{Host: host, Username: username, Password: password}
	connection.Client = new(http.Client)

	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", connection.Host, "/_api/version?details=true"), nil)

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(connection.Username, connection.Password)

	res, err := connection.Client.Do(req)

	if err != nil {
		panic("failed to establish connection")
	}

	results := util.JSONDecodeHTTPResponse(res).(map[string]interface{})

	errorNum, hasKey := results["errorNum"]
	if hasKey == true {
		errorString, _ := fmt.Printf("Failed to establish connection Host:%s ErrorNum:%g ErrorMsg:%s ", connection.Host, errorNum, results["errorMessage"])
		return nil, errors.New(string(errorString))
	}

	return connection, nil
}

// Execute will run abitary command on current connection.
func (c *Connection) Execute(httpMethod, endpoint string, statement map[string]interface{}) (results map[string]interface{}, err error) {
	compliedStatement, err := json.Marshal(statement)
	requestEndPoint := fmt.Sprintf("%s%s", c.Host, endpoint)

	req, err := http.NewRequest(httpMethod, requestEndPoint, bytes.NewBuffer(compliedStatement))

	if err != nil {
		return nil, errors.New("Failed to create a request : " + err.Error())
	}

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(c.Username, c.Password)

	log.Printf("Executing query EndPoint:%s Host:%s Statement:%s ", requestEndPoint, c.Host, compliedStatement)

	res, err := c.Client.Do(req)

	if err != nil {
		return nil, errors.New("Failed to run request : " + err.Error())
	}

	results = util.JSONDecodeHTTPResponse(res).(map[string]interface{})

	errorNum, hasKey := results["errorNum"]
	if hasKey == true {
		errorString, _ := fmt.Printf("Failed to execute query EndPoint:%s Host:%s ErrorNum:%g ErrorMsg:%s Statement:%s ", endpoint, c.Host, errorNum, results["errorMessage"], compliedStatement)
		return nil, errors.New(string(errorString))
	}

	return results, nil
}
