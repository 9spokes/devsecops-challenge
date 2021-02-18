package main

import (
	"fmt"
	"net/http"
)

const port = 8080

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
	})
	fmt.Printf("Starting hello server on port %d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
