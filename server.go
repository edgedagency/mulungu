package mulungu

import (
	"net/http"

	"github.com/gorilla/mux"
	"golang.org/x/net/context"
)

//Server server interface
type Server interface {
	Start()
	RegisterHandler(path string, f func(context.Context, http.ResponseWriter, *http.Request)) *mux.Route
}
