package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	//Here, we form a new HTTP request. This is the request that's going to be
	// passed to our handler.
	// The first argument is the method, the second argument is the route (which 
	//we leave blank for now, and will get back to soon), and the third is the 
	//request body, which we don't have in this case.
	req, err := http.NewRequest("GET", "/health", nil)

	// In case there is an error in forming the request, we fail and stop the test
	if err != nil {
		t.Fatal(err)
	}

	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

    if body := response.Body.String(); body != "Hello World!" {
        t.Errorf("Expected a hello world. Got %s", body)
    }
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
    rr := httptest.NewRecorder()
    InitRouter().ServeHTTP(rr, req)

    return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
    if expected != actual {
        t.Errorf("Expected response code %d. Got %d\n", expected, actual)
    }
}