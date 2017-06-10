package mulungu

import "encoding/json"

// //Response data structured used to communicate responses back to client
// type Response struct {
// 	Data    interface{} `json:"data"`
// 	Message string      `json:"message"`
// 	Error   bool        `json:"error"`
// }
//
// //NewResponse web function used to create new response
// func NewResponse(data interface{}, message string, err bool) *Response {
// 	return &Response{Data: data, Message: message, Error: err}
// }
//
// func (r *Response) JSON []byte{
//   jData, err := json.Marshal(Data)
// }
//
// //ToMap convert to map
// func (r *Response) ToMap() map[string]interface{} {
// 	return structs.Map(r)
// }

//Response response object
type Response map[string]interface{}

//NewResponse create a new response
func NewResponse() Response {
	return make(Response)
}

//Add add value to response map object
func (r Response) Add(key string, value interface{}) Response {
	r[key] = value

	return r
}

//Get add value to response map object
func (r Response) Get(key string) interface{} {
	return r[key]
}

//JSON converts to json presentation
func (r Response) JSON() []byte {
	data, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}
	return data
}
