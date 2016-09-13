package web

import "github.com/gorilla/sessions"

//Session will be used for session management and control
type Session struct {
	cookieStore *sessions.CookieStore
}
