package web

import (
	"net/http"

	"github.com/gorilla/sessions"
)

const (
	authorization = "Authorization"
)

//Middleware interface
type Middleware func(w http.ResponseWriter, r *http.Request, sessions *sessions.CookieStore)

//MiddlewareBasicAuthentication interface method
func MiddlewareBasicAuthentication(w http.ResponseWriter, r *http.Request, sessions *sessions.CookieStore) {

	// if len(r.Header.Get(authorization)) > 0 {
	// 	basic := r.Header.Get(authorization)
	// 	authenticationParts := strings.Split(basic, " ")
	// 	username, password, err := util.BasicAuthUsernamePassword(authenticationParts[1])
	//
	// 	if err != nil {
	// 		log.Fatal("Unable to process string")
	// 	}
	//
	// 	fmt.Printf("username: %s password: %s", username, password)
	// }
	//
	// defer r.Body.Close()
	// //TODO: setup URLS that need to be skipped to allow processing of user credentials
	// //is user authenticated?
	// //sessions.Get(r, authenticateUser)
	// b, err := ioutil.ReadAll(r.Body)
	// if err != nil {
	// 	w.WriteHeader(http.StatusExpectationFailed)
	// 	w.Write(nil)
	// }
	//
	// fmt.Println(b)
	//
	// w.Header().Add("WWW-Authenticate", "Basic realm=\"Secured Realm\"")
	// w.WriteHeader(http.StatusUnauthorized)
	// w.Write(nil)
	//io.WriteString(w, "Authentication required to access these resources")
}
