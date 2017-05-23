package middleware

import (
	"net/http"

	"golang.org/x/net/context"

	"github.com/edgedagency/mulungu/configuration"
	"github.com/edgedagency/mulungu/constant"

	"google.golang.org/appengine"
)

//type key int

const requestIDKey key = 0

//Configuration this middleware will load configurations and pass on to rest of request
func Configuration(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := appengine.NewContext(r.Context())

		//READ THIS: https://www.nicolasmerouze.com/share-values-between-middlewares-context-golang/
		configuration.NewConfigurationEntryModel(ctx, r.Header.Get(constant.HeaderNamespace))
		//LETS CONSIDER SOMTHING LIKE this

		// ctx.Value(key) <=== Where does it get it's value from based on key, how do we populate the ctx

		next.ServeHTTP(w, r)
	})
}

//NewContextWithRequestID populates context so that for each request theres an ID associated with it
func NewContextWithRequestID(ctx context.Context, r *http.Request, key string) context.Context {
	reqID := r.Header.Get(key) //Gets "X-Request-ID from HTTP Header "

	return context.WithValue(ctx, requestIDKey, reqID)

}

//ConfigurationMiddleware create a new Request value from context, and pass it onto the next handler .
func ConfigurationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ctx := NewContextWithRequestID(r.Context(), r)
		configuration.NewConfigurationEntryModel(ctx, r.Header.Get(constant.HeaderNamespace))

		next.ServeHTTP(w, r)
	})
}

//RequestIDFromContext populates its context with value key
func RequestIDFromContext(ctx context.Context, key string) string {
	return ctx.Value(key)
}

func handler(w http.ResponseWriter, r *http.Request) {
	reqID := RequestIDFromContext(r.Context())
	return w, reqID
}
