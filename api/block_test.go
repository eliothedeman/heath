package api_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetBlock(t *testing.T) {
	_, _, e := harness(t)
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/api/block", nil)
	e.ServeHTTP(w, r)

	if w.Code != 400 {
		t.Errorf("Wanted 400 got %d", w.Code)
	}
}
