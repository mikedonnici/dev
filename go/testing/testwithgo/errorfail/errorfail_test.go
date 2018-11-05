package errorfail

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {

	// First set are critical so t.Fail() is appropriate

	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "", nil)
	if err != nil {
		t.Fatalf("http.NewRequest() err = %s", err)
	}

	Handler(w, r)
	resp := w.Result()
	if resp.StatusCode != 200  {
		t.Fatalf("Want stauc code 200, got %v", resp.StatusCode)
	}

	xb, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Could not read body - %s", err)
	}

	var p Person
	err = json.Unmarshal(xb, &p)
	if err != nil {
		t.Fatalf(" json.NewDecoder() err = %s", err)
	}

	// Form here, t.Error() is ok

	want := "Michael"
	got := p.FirstName
	if got != want {
		t.Errorf("FirstName: want %q, got %q", want, got)
	}

	want = "Donnici"
	got = p.LastName
	if got != want {
		t.Errorf("LastName: want %q, got %q", want, got)
	}

	wantAge := 47
	gotAge := p.Age
	if gotAge != wantAge {
		t.Errorf("Age: want %d, got %d", wantAge, gotAge)
	}
}
