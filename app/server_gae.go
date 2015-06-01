// +build appengine

package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

func init() {

	usvc := NewUserService()

	// load all the handers for the appengine project
	r := mux.NewRouter()
	r.HandleFunc("/users", AnonHandler(usvc.NewHandler)).Methods("POST")
	http.Handle("/", r)
}

// newContext returns a context of the in-flight request r.
func newContext(r *http.Request) context.Context {
	return appengine.NewContext(r)
}

// logf logs an info message using appengine's context.
func logf(c context.Context, format string, args ...interface{}) {
	log.Infof(c, format, args...)
}

// errorf logs an error message using appengine's context.
func errorf(c context.Context, format string, args ...interface{}) {
	log.Errorf(c, format, args...)
}
