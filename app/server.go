// +build !appengine

package main

import (
	"log"
	"net/http"
	"os"

	"golang.org/x/net/context"
)

func main() {

	usvc := NewUserService()

	// load all the handers for the appengine project
	//	r := mux.NewRouter()
	//	r.HandleFunc("/users", AnonHandler(usvc.NewHandler))
	http.HandleFunc("/users", AnonHandler(usvc.NewHandler))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		// don't need context here
		errorf(nil, "%v", err)
		os.Exit(1)
	}

}

// newContext returns a context of the in-flight request r.
func newContext(r *http.Request) context.Context {
	return context.Background()
}

// logf logs an info message using Go's standard log package.
func logf(_ context.Context, format string, args ...interface{}) {
	log.Printf(format, args...)
}

// errorf logs an error message using Go's standard log package.
func errorf(_ context.Context, format string, args ...interface{}) {
	log.Printf("ERROR: "+format, args...)
}
