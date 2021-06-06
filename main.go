package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/process", func(rw http.ResponseWriter, r *http.Request) {
		// r.ParseForm()
		fmt.Fprintln(rw, r.FormValue("first_name"))
	})

	http.ListenAndServe(":8080", nil)
}
