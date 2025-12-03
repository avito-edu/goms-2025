// server.go
package main

import (
	"net/http"
)

func HelloHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	w.Write([]byte("Hello, World!"))
}

func main() {
	http.HandleFunc("/", HelloHandler)
	http.ListenAndServe(":8080", nil)
}
