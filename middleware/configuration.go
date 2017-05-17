package middleware

import (
	"net/http"

	"github.com/edgedagency/mulungu/configuration"
	"github.com/edgedagency/mulungu/constant"

	"google.golang.org/appengine"
)

//Configuration this middleware will load configurations and pass on to rest of request
func Configuration(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := appengine.NewContext(r)

		//READ THIS: https://www.nicolasmerouze.com/share-values-between-middlewares-context-golang/
		configuration.NewConfigurationEntryModel(ctx, r.Header.Get(constant.HeaderNamespace))
		//LETS CONSIDER SOMTHING LIKE this
		// ctx.Value(key) <=== Where does it get it's value from based on key, how do we populate the ctx

		next.ServeHTTP(w, r)
	})
}
