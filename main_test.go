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

func TestIncreasePepper(t *testing.T) {
	want := GetScore("Pepper")
	request, _ := http.NewRequest(http.MethodPost, "/players/Pepper", nil)
	response := httptest.NewRecorder()

	PlayerServer(response, request)

	got := GetScore("Pepper")

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

// this test has to fail since we are testing a random value
// func TestGETRandom(t *testing.T) {
// 	request, _ := http.NewRequest(http.MethodGet, "/players/beach", nil)
// 	response := httptest.NewRecorder()

// 	PlayerServer(response, request)

// 	got := response.Body.String()
// 	want := "0"

//		if got != want {
//			t.Errorf("got %q, want %q", got, want)
//		}
//	}
func TestPlayers(t *testing.T) {

	gotPepper := GetScore("Pepper")
	wantPepper := 20
	if gotPepper != wantPepper {
		t.Errorf("got %d, want %d", gotPepper, wantPepper)
	}

	gotSalt := GetScore("Salt")
	wantSalt := 10
	if gotSalt != wantSalt {
		t.Errorf("got %d, want %d", gotSalt, wantSalt)
	}

	gotPaprika := GetScore("Paprika")
	wantPaprika := 30
	if gotPaprika != wantPaprika {
		t.Errorf("got %d, want %d", gotPaprika, wantPaprika)
	}
	gotCurry := GetScore("Curry")
	wantCurry := 0
	if gotCurry != wantCurry {
		t.Errorf("got %d, want %d", gotCurry, wantCurry)
	}
}
func TestIncreaseScore(t *testing.T) {
	ResetScores("Pepper")
	wantPepper := 21
	gotPepper := IncreaseScore("Pepper")
	if gotPepper != wantPepper {
		t.Errorf("got %d, want %d", gotPepper, wantPepper)
	}
	wantPepper = 22
	gotPepper = IncreaseScore("Pepper")
	if gotPepper != wantPepper {
		t.Errorf("got %d, want %d", gotPepper, wantPepper)
	}

}

func TestResetScores(t *testing.T) {
	wantPepper := 20
	gotPepper := ResetScores("Pepper")
	if gotPepper != wantPepper {
		t.Errorf("got %d, want %d", gotPepper, wantPepper)
	}
	gotPaprika := ResetScores("Paprika")
	wantPaprika := 30
	if gotPaprika != wantPaprika {
		t.Errorf("got %d, want %d", gotPaprika, wantPaprika)
	}
}
