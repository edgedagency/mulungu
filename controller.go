package mulungu

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/edgedagency/mulungu/configuration"
	"github.com/edgedagency/mulungu/util"
	"golang.org/x/net/context"
	"google.golang.org/appengine/log"
)

//Controller provides basic controller functionionlity
type Controller struct {
	Config *configuration.Config
}

//Data returns request body as map[string]interface{}
func (c *Controller) Data(ctx context.Context, w http.ResponseWriter, r *http.Request) map[string]interface{} {
	data, err := util.JSONDecodeHTTPRequest(r)
	if err != nil {
		log.Errorf(ctx, "failed to decode request, error %s", err.Error())
		c.JSONResponse(w, NewResponse(map[string]interface{}{"message": "unable to decode request", "error": err.Error()},
			"failed to process request", true), http.StatusBadRequest)
		return nil
	}
	return data
}

//HydrateModel hydrates model from request body
func (c *Controller) HydrateModel(ctx context.Context, readCloser io.ReadCloser, dest interface{}) error {
	err := json.NewDecoder(readCloser).Decode(dest)
	if err != nil {
		log.Errorf(ctx, "failed to hydrate model, %s", err.Error())
		return err
	}
	return nil
}

//JSONResponse returns json response and sets right content headers
func (c *Controller) JSONResponse(w http.ResponseWriter, responseBody interface{}, statusCode int) {
	util.WriteJSON(w, responseBody, statusCode)
}
