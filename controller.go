package mulungu

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/edgedagency/mulungu/util"
	"github.com/gorilla/mux"
	"golang.org/x/net/context"
	"google.golang.org/appengine/log"
)

//Controller provides basic controller functionionlity
type Controller struct {
	//Config *configuration.Config
}

//Data returns request body as map[string]interface{}
func (c *Controller) Data(ctx context.Context, w http.ResponseWriter, r *http.Request) map[string]interface{} {
	data, err := util.JSONDecodeHTTPRequest(r)
	if err != nil {
		log.Errorf(ctx, "failed to decode request, error %s", err.Error())
		// c.JSONResponse(w, NewResponse(map[string]interface{}{"message": "unable to decode request", "error": err.Error()},
		// 	"failed to process request", true), http.StatusBadRequest)
		return nil
	}
	return data
}

//PathValue obtians value from path variable configurations
func (c *Controller) PathValue(r *http.Request, key, defaultValue string) string {
	pathValues := mux.Vars(r)
	if value, ok := pathValues[key]; ok {
		return value
	}
	return defaultValue
}

//ParamValue obtains param value from url ?env=dev&expire-date=1896
func (c *Controller) ParamValue(r *http.Request, key, defaultValue string) string {
	paramValue := r.FormValue(key)
	if len(paramValue) > 0 {
		return paramValue
	}
	return defaultValue
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

//WriteJSON respond to request
func (c *Controller) WriteJSON(w http.ResponseWriter, statusCode int, bytes []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(bytes)
}

//WriteXML respond to request
func (c *Controller) WriteXML(w http.ResponseWriter, statusCode int, bytes []byte) {
	w.Header().Set("Content-Type", "application/xml")
	w.WriteHeader(statusCode)
	w.Write(bytes)
}

//ResponseBodyToBytes obtains response body as bytes or empty
func (c *Controller) ResponseBodyToBytes(r *http.Response) []byte {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return []byte{}
	}

	return bytes
}
