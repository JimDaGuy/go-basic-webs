package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hiya")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Requested: %s\n", r.URL.Path)
	})

	http.ListenAndServe(":8080", nil)
}
