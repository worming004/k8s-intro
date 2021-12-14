package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

type backendProxy struct {
}

func (bp backendProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	res, err := http.Get("http://http-backend/hello")
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	fmt.Println("here is the gateway reaching the backend")
	w.Write(body)
}

func main() {
	mux := mux.NewRouter()
	mux.Handle("/proxy", backendProxy{}).Methods("GET")
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	}).Methods("GET")

	http.ListenAndServe(":8080", mux)
}
