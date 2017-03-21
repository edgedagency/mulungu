package middleware

import (
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

//Logging this handler determins and set appropriate namespace on re
func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := appengine.NewContext(r)

		log.Debugf(ctx, "**MIDDLEWARE** logging request information")

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

		next.ServeHTTP(w, r)

	})
}
