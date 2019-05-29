package utils

import (
	"log"
	"net/http"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
)

// var store = ResisSesssion()

var store *sessions.CookieStore = sessions.NewCookieStore(securecookie.GenerateRandomKey(64))

// HandlerFunc is a sugar syntax for applicationHandler
func HandlerFunc(r *mux.Router, path string, fn http.HandlerFunc) *mux.Route {
	return r.HandleFunc(path, applicationHandler(fn))
}

// applicationHandler is a common function among all handlers.
func applicationHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Get Session
		session, err := store.Get(r, SessionName)
		if err != nil {
			// Recreate session
			session, err := store.New(r, SessionName)
			if err != nil {
				log.Fatal(err)
			}

			context.Set(r, ContextSessionKey, session)
			// Invoke handler
			fn(w, r)
		}
		context.Set(r, ContextSessionKey, session)
		// Invoke handler
		fn(w, r)
	}
}
