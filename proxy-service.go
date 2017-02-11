package mulungu

import (
	"net/http"

	"github.com/edgedagency/mulungu/util"
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	log "google.golang.org/appengine/log"
)

//ProxyService object used to make proxy services
type ProxyService struct {
	Context context.Context
	Schema  string
	Service string
	Path    string
}

//NewProxyService returns proxy service handler
func NewProxyService() *ProxyService {
	proxyService := &ProxyService{}
	return proxyService
}

//Send used to send proxy request
func (ps *ProxyService) Send(method string, data map[string]interface{}, username, password string, secured bool, headers map[string]string) (*HTTPResponse, error) {
	serviceHost := GenerateGoogleServiceHost(appengine.DefaultVersionHostname(ps.Context), ps.Service)
	httpResponse := NewHTTPRequest(ps.Context, ps.Schema, serviceHost, username, password, secured, headers).SendJSON(method, ps.Path, data)

	if httpResponse.HasErrors() {
		log.Errorf(ps.Context, "failed to processes/proxy request, error: %s", httpResponse.Error.Error())
		return httpResponse, httpResponse.Error
	}

	return httpResponse, nil
}

//Proxy shortcut method that executes a proxy request
func (ps *ProxyService) Proxy(w http.ResponseWriter, r *http.Request) (*HTTPResponse, error) {
	data, dataErr := util.JSONDecodeHTTPRequest(r)
	if dataErr != nil {
		log.Errorf(ps.Context, "failed to decode body, error : %s", dataErr.Error())
	}
	log.Debugf(ps.Context, "firing off request, data:%v service:%s path:%s", data, ps.Service, ps.Path)
	return ps.Send(r.Method, data, "", "", false, map[string]string{"Content-Type": "application/json", "X-Namespace": r.Header.Get("Host")})
}

//SetUsername sets username to use for this proxy service
func (ps *ProxyService) SetUsername(username string) {

}
