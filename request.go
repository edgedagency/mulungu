package mulungu

import (
	"net/http"

	"golang.org/x/net/context"
)

//Request object that will be passed to represent request received
type Request struct {
	CTX  context.Context
	Data *RequestData
}

//NewRequest creates a new request object
func NewRequest(ctx context.Context, httpRequest *http.Request) *Request {
	r := &Request{}
	r.CTX = ctx
	r.Data = NewRequestData(httpRequest)

	return r
}
