package mulungu

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

//AppEngineServer used when application is deployed to google app server
type AppEngineServer struct {
	router *mux.Router
}

//NewAppEngineServer creates server which can function on GAE servers
func NewAppEngineServer() *AppEngineServer {
	aes := AppEngineServer{router: mux.NewRouter()}
	return &aes
}

//Start starts server services
func (aes *AppEngineServer) Start() {
	http.Handle("/", aes.router)
}

//Router gets server router
func (aes *AppEngineServer) Router() *mux.Router {
	return aes.router
}

//RegisterHandler registers handlers
func (aes *AppEngineServer) RegisterHandler(path string, f func(ctx context.Context, w http.ResponseWriter, r *http.Request)) *mux.Route {
	return aes.router.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		//determine this hosts configurations and state before processing request. Request filter e.t.c.
		ctx, ctxErr := aes.GetContext(r)

		if ctxErr != nil {
			ctx = appengine.NewContext(r)
			log.Errorf(ctx, "Failed to obtain custom namespace context, error:%s", ctxErr.Error())
		}

		log.Debugf(ctx, "Request Details")

		log.Debugf(ctx, "URI:%s DefaultHostName:%s Host:%s Path:%s Referer:%s IP:%s Scheme:%s Method:%s X-Namespace:%s appId:%s datacenter:%s environment:%s",
			r.URL.RequestURI(),
			appengine.DefaultVersionHostname(ctx),
			r.URL.Host,
			r.URL.Path,
			r.Referer(),
			r.RemoteAddr,
			r.URL.Scheme,
			r.Method,
			r.Header.Get("X-Namespace"),
			appengine.AppID(ctx),
			appengine.Datacenter(ctx),
			r.Header.Get("X-Environment"))

		log.Debugf(ctx, "Headers")
		for k, v := range r.Header {
			log.Debugf(ctx, "key:%s value:%s", k, v)
		}
		log.Debugf(ctx, "Host Information")
		log.Debugf(ctx, "Zoo %s", r.Header.Get("X-Zoo"))
		zooParts := strings.Split(r.Header.Get("X-Zoo"), ",")
		for _, partValue := range zooParts {
			log.Debugf(ctx, "Zoo Part %s", partValue)
		}

		// w.Header().Set("Access-Control-Allow-Origin", "*")
		// w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		//
		// if r.Method == http.MethodOptions {
		// 	w.Header().Set("Access-Control-Allow-Origin", "*")
		// 	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		// 	return
		// }

		f(ctx, w, r)
	})
}

//GetContext returns new context based on host
func (aes *AppEngineServer) GetContext(r *http.Request) (context.Context, error) {
	ctx := appengine.NewContext(r)
	aes.SetNamespace(ctx, r)
	aes.SetEnvironmentOnNamespace(ctx, r)

	return appengine.Namespace(ctx, r.Header.Get("X-Namespace"))
}

//SetNamespace sets environment to be used
func (aes *AppEngineServer) SetNamespace(ctx context.Context, r *http.Request) {
	if r.Header.Get("X-Namespace") == "" {
		if r.Header.Get("Host") != "" {
			log.Debugf(ctx, "using Host as namespace:%s", r.Header.Get("Host"))
			r.Header.Set("X-Namespace", r.Header.Get("Host"))
		} else if r.URL.Host != "" {
			log.Debugf(ctx, "using URL Host as namespace:%s", r.URL.Host)
			r.Header.Set("X-Namespace", r.URL.Host)
		}
	}
}

//SetEnvironmentOnNamespace sets sets namespace environmnet e.g. dev.namespace.xyz
func (aes *AppEngineServer) SetEnvironmentOnNamespace(ctx context.Context, r *http.Request) {
	environment := r.URL.Query().Get("env")
	if environment != "" {
		log.Debugf(ctx, "setting environment on environment:%s", environment)
		r.Header.Set("X-Environment", environment)
		r.Header.Set("X-Namespace", strings.Join([]string{r.Header.Get("X-Environment"), r.Header.Get("X-Namespace")}, "."))
		log.Debugf(ctx, "environment namespace:%s", r.Header.Get("X-Namespace"))
	} else {
		log.Debugf(ctx, "no environment specified for request, use http://example.com?env=dev to set environment to dev")
	}
}
