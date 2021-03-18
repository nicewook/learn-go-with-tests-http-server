package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETPlayers(t *testing.T) {
	t.Run("returns Pepper's score", func(t *testing.T) {
		req, err := NewGetScoreRequest("Pepper")
		if err != nil {
			t.Error(err)
		}
		resp := httptest.NewRecorder()

		PlayerServer(resp, req)

		assertResponseBody(t, resp.Body.String(), "20")
	})

	t.Run("returns Floyd's score", func(t *testing.T) {
		req, err := NewGetScoreRequest("Floyd")
		if err != nil {
			t.Error(err)
		}
		resp := httptest.NewRecorder()

		PlayerServer(resp, req)

		assertResponseBody(t, resp.Body.String(), "10")
	})
}

// Helper function
func NewGetScoreRequest(name string) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodGet, "/players/"+name, nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func assertResponseBody(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
