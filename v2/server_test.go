package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETPlayers(t *testing.T) {
	t.Run("returns Pepper's score", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "player/Pepper", nil)
		if err != nil {
			t.Error(err)
		}
		resp := httptest.NewRecorder()

		PlayerServer(resp, req)

		got := resp.Body.String()
		want := "20"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
