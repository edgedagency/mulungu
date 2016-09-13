package mulungu

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/edgedagency/mulungu/util"
)

// Connection is a structure which will be used to represent a connection
type Connection struct {
	host     string
	username string
	password string
	client   *http.Client
}

// NewConnection creates connection and returns &pointer to refrence to this connection.
func NewConnection(host, username, password string) *Connection {
	return &Connection{host: host, username: username, password: password, client: new(http.Client)}
}

// Execute will run abitary command on current connection.
func (c *Connection) Execute(httpMethod, endpoint string, statement map[string]interface{}) (response map[string]interface{}, err error) {
	compliedStatement, err := json.Marshal(statement)
	requestEndPoint := fmt.Sprintf("%s%s", c.host, endpoint)

	fmt.Println(fmt.Sprintf("statement %s endpoint %s ", compliedStatement, requestEndPoint))
	req, err := http.NewRequest(httpMethod, requestEndPoint, bytes.NewBuffer(compliedStatement))

	if err != nil {
		return nil, fmt.Errorf("failed to create http request request error %s", err.Error())
	}

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(c.username, c.password)

	res, err := c.client.Do(req)

	if err != nil {
		return nil, err
	}

	fmt.Println("Connection Response")
	fmt.Println(res)

	return util.JSONDecodeHTTPResponse(res)
}
