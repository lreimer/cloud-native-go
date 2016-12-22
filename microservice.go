package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/api/hello", hello)
	http.HandleFunc("/api/echo", echo)

	http.ListenAndServe(":8080", nil)
}

// Hello response structure
type Hello struct {
	Message string
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Cloud Native Go.")
}

func echo(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Query()["message"][0]

	w.Header().Add("Content-Type", "text/plain")
	fmt.Fprintf(w, message)
}

func hello(w http.ResponseWriter, r *http.Request) {

	m := Hello{"Welcome to Cloud Native Go."}
	b, err := json.Marshal(m)

	if err != nil {
		panic(err)
	}

	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Write(b)
}
