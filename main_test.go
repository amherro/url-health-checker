package main

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestURLChecker(t *testing.T) {
	t.Run("Test two successful URLs", func(t *testing.T) {
		serverOne := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}))

		serverTwo := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}))

		defer serverOne.Close()
		defer serverTwo.Close()

		urls := []string{serverOne.URL, serverTwo.URL}

		got := URLChecker(urls)
		want := map[string]int{serverOne.URL: 200, serverTwo.URL: 200}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got %v want %v", got, want)
		}
	})
	t.Run("Test a URL that was not found", func(t *testing.T) {
		server404 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
		}))
		serverSuccess := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}))

		defer server404.Close()
		defer serverSuccess.Close()

		urls := []string{server404.URL, serverSuccess.URL}

		got := URLChecker(urls)
		want := map[string]int{server404.URL: 404, serverSuccess.URL: 200}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got %v want %v", got, want)
		}
	})
	t.Run("Test a DOMAIN that doesn't exist", func(t *testing.T) {
		serverExists := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}))

		defer serverExists.Close()

		urls := []string{"http://127.0.0.1:1", serverExists.URL}
		got := URLChecker(urls)
		want := map[string]int{"http://127.0.0.1:1": 0, serverExists.URL: 200}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got %v want %v", got, want)
		}
	})
}
