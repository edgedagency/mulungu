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
	Server
	Config *configuration.Config
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
