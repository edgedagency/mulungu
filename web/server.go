package web

import (
	"fmt"
	"net/http"
	"os"

	"github.com/edgedagency/mulungu"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

//Server used to manage and process server requests
type Server struct {
	router        *mux.Router
	middleware    map[string]Middleware
	sessions      *sessions.CookieStore
	configuration *mulungu.Configuration
}

//NewServer starts up servers and services
func NewServer() *Server {
	//TODO implement an event based pattern which will annocue event like server booted enabling decoubled mutation and access to this server
	server := &Server{router: mux.NewRouter(), middleware: map[string]Middleware{},
		sessions: sessions.NewCookieStore([]byte("something-very-secret")), configuration: mulungu.NewConfiguration(os.Getenv("SPZA_BASE_PATH"))}
	return server
}

//Boot boots up server
func (s *Server) Boot() {
	http.ListenAndServe(":9001", context.ClearHandler(s))
	//TODO read server address from environmental variable, if that fails randomize server port between 4000 to 4999

}

//ServeHTTP
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("-----------------------------")
	fmt.Println(fmt.Sprintf("Received request from %s to path %s ", r.RemoteAddr, r.RequestURI))
	s.router.ServeHTTP(w, r)
}

//SetMiddleware addds middleware handler to server
func (s *Server) SetMiddleware(name string, middleware Middleware) {
	s.middleware[name] = middleware
}

//SetRoute method used to register routes
func (s *Server) SetRoute(path string, handler http.HandlerFunc) *mux.Route {
	return s.router.HandleFunc(path, handler)
}
