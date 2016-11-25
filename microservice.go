package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/api/hello", hello)
	http.ListenAndServe(":8080", nil)
}

// Hello response structure
type Hello struct {
	Message string
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Cloud Native Go.")
}

func hello(w http.ResponseWriter, r *http.Request) {

	m := Hello{"Welcome to Cloud Native Go with Wercker on DC/OS. Hello Fuchsi und Dominik."}
	b, err := json.Marshal(m)

	if err != nil {
		panic(err)
	}

	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Write(b)
}
