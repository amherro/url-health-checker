package main

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestURLChecker(t *testing.T) {
	newServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer newServer.Close()

	urls := []string{"https://adamherro.dev/", "https://vaudify.com/"}

	got := URLChecker(urls)
	want := map[string]bool{"https://adamherro.dev/": true, "https://vaudify.com/": true}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got %v want %v", got, want)
	}
}
