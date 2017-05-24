package middleware

import (
	"net/http"

	"google.golang.org/appengine"

	"github.com/edgedagency/mulungu/constant"
	"github.com/edgedagency/mulungu/logger"
)

//ConfigurationMiddleware create a new Request value from context, and pass it onto the next handler .
func ConfigurationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := appengine.NewContext(r)

		// 1. obtain namespace from header
		namespace := r.Header.Get(constant.HeaderNamespace)
		logger.Debugf(ctx, "middleware configurations", "getting configurations with namespace %s", namespace)

		// 2. call model which retireves all configurations entries
		// configurationsModel := configuration.NewConfigurationEntryModel(ctx, namespace)

		// 3. execute findAll on model to return entries from datastore
		// configurations := configurationsModel.findAll()
		// logger.Debugf(ctx, "middleware configurations", "getting configurations with namespace %#v", configurations)

		next.ServeHTTP(w, r)

	})
}
