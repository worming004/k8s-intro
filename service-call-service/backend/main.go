package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	mux := mux.NewRouter()

	mux.HandleFunc("/hello", helloWorldHandlerFunc).Methods("GET")

	port := ":8080"
	fmt.Println("listen on port " + port)
	http.ListenAndServe(port, mux)
}

func helloWorldHandlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world\n"))
}
