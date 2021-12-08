package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthCheck(t *testing.T) {
	handlerBuilder := newHealthHandler()

	isHealthyHandler := handlerBuilder.buildHealthHandler()
	setUnhealthyHandler := handlerBuilder.setUnhealthyHandler()

	recorder := httptest.NewRecorder()
	isHealthyHandler.ServeHTTP(recorder, nil)

	if recorder.Code != http.StatusOK {
		t.Errorf("expecting default behavior as healthy")
	}

	recorder = httptest.NewRecorder()
	setUnhealthyHandler(recorder, nil)

	recorder = httptest.NewRecorder()
	isHealthyHandler.ServeHTTP(recorder, nil)

	if recorder.Code == http.StatusOK {
		t.Errorf("expecting unhealthy after set")
	}
}
