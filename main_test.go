package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRouter(t *testing.T) {
	r := newRouter()
	mockServer := httptest.NewServer(r)

	resp, err := http.Get(mockServer.URL + "/assets/view")

	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status should be ok, got %d", resp.StatusCode)
	}

	defer resp.Body.Close()

}