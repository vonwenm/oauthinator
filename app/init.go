// +build appengine

package oauthinator

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

type appHandler func(*http.Request) (interface{}, error)

func init() {

	usvc := NewUserService()

	// load all the handers for the appengine project
	r := mux.NewRouter()
	r.HandleFunc("/users", anonHander(usvc.NewHandler)).Methods("POST")
	http.Handle("/", r)
}

func anonHander(h appHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := appengine.NewContext(r)

		resp, err := h(r)

		// build wrapped json response
		appResp := &appResponse{Response: resp}

		if err != nil {
			log.Errorf(c, "%v", err)
			appResp.Error = err.Error()
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		if err = json.NewEncoder(w).Encode(appResp); err != nil {
			log.Criticalf(c, "Error encoding response: %v", err)
		}
	}
}
