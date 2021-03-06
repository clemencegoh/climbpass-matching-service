package main

import (
	"climbpass-matching-service/constants"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Integration Test Suite
func TestHealthCheck(t *testing.T) {
	//Here, we form a new HTTP request. This is the request that's going to be
	// passed to our handler.
	req, err := http.NewRequest("GET", constants.APIBasePath+"/health", nil)

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

func TestGymAPI(t *testing.T) {
	req, err := http.NewRequest("GET", constants.APIBasePath+"/gyms", nil)
	if err != nil {
		t.Fatal(err)
	}

	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)
	if body := response.Body.String(); body != "[]" {
		t.Errorf("Expected an empty array. Got %s", body)
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	InitRouter("test").ServeHTTP(rr, req)
	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
