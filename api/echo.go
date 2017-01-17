package api

import (
	"fmt"
	"net/http"
)

// EchoHandleFunc to be used as http.HandleFunc for ECHO API
func EchoHandleFunc(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Query()["message"][0]

	w.Header().Add("Content-Type", "text/plain")
	fmt.Fprintf(w, message)
}
