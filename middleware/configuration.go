package middleware

import (
	"net/http"

	"google.golang.org/appengine"

	"github.com/edgedagency/mulungu/configuration"
	"github.com/edgedagency/mulungu/constant"
	"github.com/edgedagency/mulungu/logger"
	"github.com/edgedagency/mulungu/util"
)

//Configuration create a new Request value from context, and pass it onto the next handler .
func Configuration(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//return http.HandleFunc(func(filter string, w http.ResponseWriter, r *http.Request) {
		ctx := appengine.NewContext(r)

		//obtain namespace from header
		namespace := r.Header.Get(constant.HeaderNamespace)
		logger.Debugf(ctx, "middleware configurations", "getting configurations with namespace %s", namespace)

		//call model which retireves all configurations entries
		configurationsModel := configuration.NewConfigurationEntryModel(ctx, namespace)
		//execute findAll on model to return entries from datastore
		configurations, findErr := configurationsModel.FindAll(r.URL.Query().Get("filter"))
		if findErr != nil {
			logger.Errorf(ctx, "middleware configuration", "failed to find record, error:%s", findErr.Error())
		}
		logger.Debugf(ctx, "middleware configurations", "getting configurations with namespace %#v", configurations)

		//lets loop through obtained configurations
		for _, configurationEntry := range configurations {
			util.ContextSetValue(r, configurationEntry.Key, configurationEntry.value)
			logger.Debugf(ctx, "middleware configurations", "obtaining configuration key:%s value:%s", configurationEntry.Key, configurationEntry.Value)
		}

		// everything above here is called before a request
		next.ServeHTTP(w, r)
		// everything below here is called after a request
	})

}
