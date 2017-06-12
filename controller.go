package mulungu

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/edgedagency/mulungu/constant"
	"github.com/edgedagency/mulungu/logger"
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
		logger.Errorf(ctx, "Base Controller", "failed to decode request, error %s", err.Error())
		c.Error(ctx, w, r, "Failed to decode request", err)
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

//Write respond to request
func (c *Controller) Write(ctx context.Context, w http.ResponseWriter, r *http.Request, statusCode int, bytes []byte) {
	contentType := strings.TrimSpace(strings.ToLower(r.Header.Get(constant.HeaderContentType)))
	logger.Infof(ctx, "Base Controller", "writing content type: %s content bytes: %#v", contentType, bytes)

	switch contentType {
	case "application/xml", "application/xml; charset=utf-8", "text/xml; charset=utf-8", "text/xml":
		logger.Infof(ctx, "Base Controller", "writing XML content bytes: %s", string(bytes))
		c.WriteXML(ctx, w, statusCode, bytes)
	case "application/json", "application/json; charset=utf-8":
		logger.Infof(ctx, "Base Controller", "writing JSON content bytes: %s", string(bytes))
		c.WriteJSON(ctx, w, statusCode, bytes)
	default:
		logger.Infof(ctx, "Base Controller", "writing Text content: %s", string(bytes))
		c.WriteText(ctx, w, statusCode, bytes)
	}
}

//WriteText respond to request
func (c *Controller) WriteText(ctx context.Context, w http.ResponseWriter, statusCode int, bytes []byte) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(statusCode)
	w.Write(bytes)
}

//WriteJSON respond to request
func (c *Controller) WriteJSON(ctx context.Context, w http.ResponseWriter, statusCode int, bytes []byte) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	w.Write(bytes)
}

//WriteXML respond to request
func (c *Controller) WriteXML(ctx context.Context, w http.ResponseWriter, statusCode int, bytes []byte) {
	w.Header().Set("Content-Type", "application/xml; charset=utf-8")
	w.WriteHeader(statusCode)
	w.Write(bytes)
}

//IsAuthorized determines if request is authorized in some why
func (c *Controller) IsAuthorised(ctx context.Context, r *http.Request) bool {
	return util.IsAuthorised(ctx, r)
}

//NotAuthosized creates a not authorized http response
func (c *Controller) NotAuthorized(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	c.Write(ctx, w, r, http.StatusUnauthorized, NewResponse().Add("message", "authorization required").Format(r.Header.Get(constant.HeaderContentType)))
}

//Created generates a created http response with data
func (c *Controller) Created(ctx context.Context, w http.ResponseWriter, r *http.Request, data interface{}) {
	c.Write(ctx, w, r, http.StatusCreated, NewResponse().Add("message", "Record/s created").Add("data", data).Format(r.Header.Get(constant.HeaderContentType)))
}

//Found generates a found response with data
func (c *Controller) Found(ctx context.Context, w http.ResponseWriter, r *http.Request, data interface{}) {
	c.Write(ctx, w, r, http.StatusOK, NewResponse().Add("message", "Record/s retrived").Add("data", data).Format(r.Header.Get(constant.HeaderContentType)))
}

//NotFound generates a not found http response
func (c *Controller) NotFound(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	c.Write(ctx, w, r, http.StatusNotFound, NewResponse().Add("message", "Record/s not found").Format(r.Header.Get(constant.HeaderContentType)))
}

//Error generates a error response
func (c *Controller) Error(ctx context.Context, w http.ResponseWriter, r *http.Request, message string, err error) {
	c.Write(ctx, w, r, http.StatusInternalServerError, NewResponse().Add("message", message).Add("error", err.Error()).Format(r.Header.Get(constant.HeaderContentType)))
}

//OK generates http OK response with message
func (c *Controller) OK(ctx context.Context, w http.ResponseWriter, r *http.Request, message string) {
	c.Write(ctx, w, r, http.StatusOK, NewResponse().Add("message", message).Format(r.Header.Get(constant.HeaderContentType)))
}

//ResponseBodyToBytes obtains response body as bytes or empty
func (c *Controller) ResponseBodyToBytes(r *http.Response) []byte {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return []byte{}
	}

	return bytes
}
