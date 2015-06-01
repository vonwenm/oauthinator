package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var (
	// ErrInvalidForm unable to parse the form input for a post
	ErrInvalidForm = fmt.Errorf("Invalid Input")
)

type appHandler func(*http.Request) (interface{}, error)

// AnonHandler handler for anonymous requests
func AnonHandler(h appHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := newContext(r)

		resp, err := h(r)

		// build wrapped json response
		appResp := &appResponse{Response: resp}

		if err != nil {
			errorf(c, "%v", err)
			appResp.Error = err.Error()
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		if err = json.NewEncoder(w).Encode(appResp); err != nil {
			errorf(c, "Error encoding response: %v", err)
		}
	}
}

type appResponse struct {
	Response interface{}
	Error    string
}

func responseMessage(msg string) map[string]string {
	return map[string]string{
		"msg": msg,
	}
}
