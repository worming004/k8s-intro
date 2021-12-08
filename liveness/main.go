package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	healthHandler := newHealthHandler()
	router.HandleFunc("/healthz", healthHandler.buildHealthHandler())
	router.HandleFunc("/set-healthy", healthHandler.setHealthyHandler())
	router.HandleFunc("/set-unhealthy", healthHandler.setUnhealthyHandler())

	helloHandler := stringHandler("hello world\n")
	router.HandleFunc("/hello", helloHandler)
	router.HandleFunc("/", helloHandler)

	port := ":8000"
	log.Println("listening on port " + port)
	http.ListenAndServe(port, router)
}

type healthHandler struct {
	isHealthy        bool
	healthyHandler   http.HandlerFunc
	unhealthyHandler http.HandlerFunc
}

func newHealthHandler() healthHandler {
	return healthHandler{
		isHealthy:      true,
		healthyHandler: stringHandler("ok"),
		unhealthyHandler: func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			stringHandler("nok")(w, r)
		},
	}
}

func (h *healthHandler) buildHealthHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if h.isHealthy {
			h.healthyHandler(w, r)
		} else {
			h.unhealthyHandler(w, r)
		}
	}
}

func (h *healthHandler) setHealthyHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h.isHealthy = true
		h.healthyHandler(w, r)
	}
}

func (h *healthHandler) setUnhealthyHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h.isHealthy = false
		h.healthyHandler(w, r)
	}
}

func stringHandler(value string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(value))
	}
}
