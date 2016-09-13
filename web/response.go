package web

//Response data structured used to communicate responses back to client
type Response struct {
	Code  int                     `json:"code"`
	Data  *map[string]interface{} `json:"data"`
	Error ResponseError           `json:"error"`
}

//ResponseError error attached to a response
type ResponseError struct {
	Exception error  `json:"exception"`
	Message   string `json:"message"`
}

//NewResponse web function used to create new response
func NewResponse(code int, data *map[string]interface{}, err error, message string) *Response {
	return &Response{Code: code, Data: data, Error: ResponseError{Exception: err, Message: message}}
}
