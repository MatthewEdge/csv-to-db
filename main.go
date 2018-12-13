package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/load", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "WIP")
	})

	http.ListenAndServe("0.0.0.0:8080", nil)
}
