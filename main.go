package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/home", func(rw http.ResponseWriter, r *http.Request) {
		url := r.URL
		query := url.Query()

		id := query["id"]
		fmt.Fprintln(rw, id)

		name := query.Get("name")
		fmt.Fprintln(rw, name)
	})

	http.ListenAndServe(":8080", nil)
}
