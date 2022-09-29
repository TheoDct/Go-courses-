package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETPepper(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "/players/Pepper", nil)
	response := httptest.NewRecorder()

	PlayerServer(response, request)

	got := response.Body.String()
	want := "20"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestGETSalt(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "/players/Salt", nil)
	response := httptest.NewRecorder()

	PlayerServer(response, request)

	got := response.Body.String()
	want := "10"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
func TestPlayers(t *testing.T) {

	gotPepper := GETPlayers("Pepper")
	wantPepper := 20
	if gotPepper != wantPepper {
		t.Errorf("got %q, want %q", gotPepper, wantPepper)
	}

	gotSalt := GETPlayers("Salt")
	wantSalt := 10
	if gotSalt != wantSalt {
		t.Errorf("got %q, want %q", gotSalt, wantSalt)
	}

	gotPaprika := GETPlayers("Paprika")
	wantPaprika := 30
	if gotPaprika != wantPaprika {
		t.Errorf("got %q, want %q", gotPaprika, wantPaprika)
	}
}
