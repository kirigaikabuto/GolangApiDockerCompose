package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the main tutorial about docker heheheh 1111111")
	})

	fmt.Println("staring server")
	http.ListenAndServe(":80", r)
}
