package superhandler

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_Handle(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/handle", nil)
	if err != nil {
		t.Fatal(err)
	}

	h := New(nil)
	handler := http.HandlerFunc(h.Handle)
	recorder := httptest.NewRecorder()

	handler.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %d want %d", status, http.StatusOK)
	}

	rand.Seed(1)
	expected := fmt.Sprintf("Len %d", rand.Int31n(max))

	if recorder.Body.String() != expected {
		t.Errorf("handler returned wrong response: got %s want %s", recorder.Body.String(), expected)
	}
}
