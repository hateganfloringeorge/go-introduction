package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("faster url is returned", func(t *testing.T) {
		slowServer := makeDelayedServer(2 * time.Second)
		fastServer := makeDelayedServer(0 * time.Second)
		
		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, err := Racer(slowURL, fastURL, 10 * time.Second)

		if err != nil {
			t.Fatalf("received error when not expected: %v", err)
		}

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("return error after after 10 sec timeout", func(t *testing.T) {
		server := makeDelayedServer(2 * time.Second)
		
		defer server.Close()

		_, err := ConfigurableRacer(server.URL, server.URL, 1 * time.Second)

		if err == nil {
			t.Fatal("expected error after timeout")
		}
	})
	
}

func makeDelayedServer(delay time.Duration) *httptest.Server{
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
	return server
}