package mulungu

//Response data structured used to communicate responses back to client
type Response struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Error   bool        `json:"error"`
}

//NewResponse web function used to create new response
func NewResponse(data interface{}, message string, err bool) *Response {
	return &Response{Data: data, Message: message, Error: err}
}
