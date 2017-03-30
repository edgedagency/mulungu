package util

import (
	"net/http"
	"strings"

	"github.com/edgedagency/mulungu/logger"
	"golang.org/x/net/context"
)

//IsAuthorised checks if Request is authorised
func IsAuthorised(ctx context.Context, r *http.Request) bool {
	requestAuthorised := r.Header.Get("X-AUTHORISED")
	logger.Debugf(ctx, "security util", "checking if request has been authorised, authorised %s", requestAuthorised)

	if requestAuthorised != "" && requestAuthorised == "true" {
		logger.Debugf(ctx, "security util", "request contains authorization information, authorised %s", requestAuthorised)
		return true
	}
	return false
}

//HasAnyRole checks if this request has any of the roles specified in roles
func HasAnyRole(ctx context.Context, roles []string, r *http.Request) bool {
	requestRoles := strings.Split(r.Header.Get("X-AUTHORISED-ROLES"), ",")
	logger.Debugf(ctx, "security util", "request roles %#v ", requestRoles)

	if len(requestRoles) > 0 {
		for _, role := range roles {
			for _, requestRole := range requestRoles {
				if strings.Compare(role, requestRole) == 0 {
					logger.Debugf(ctx, "security util", "request has one required role required:%s found:%s", role, requestRole)
					return true
				}
			}
		}

	}

	return false
}
