package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintln(w, "HelloWorld", r.URL.Path)
	if err != nil {
		panic(err)
	}
}

func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		panic(err)
	}
}
