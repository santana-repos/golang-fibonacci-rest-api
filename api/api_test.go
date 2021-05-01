package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthCheckHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/health_check", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	healthApiImpl := NewHealthApiImpl()
	handler := http.HandlerFunc(healthApiImpl.healthResourcesHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `Health check`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestFibonaccibHandlerRetrieve72(t *testing.T) {
	req, err := http.NewRequest("GET", "/fibonacci?n=72", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	fibApiImpl := NewFibApiImpl()
	handler := http.HandlerFunc(fibApiImpl.fibonacciResourcesHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	want := `{"fibonacciNthPosition":72,"value":"498454011879264"}`
	if rr.Body.String() != want {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), want)
	}
}

func TestFibHandlerRetrieveErrorMethodOptionsNotAllowed(t *testing.T) {
	req, err := http.NewRequest("OPTIONS", "/fib?n=72", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	fibApiImpl := NewFibApiImpl()
	handler := http.HandlerFunc(fibApiImpl.fibonacciResourcesHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	want := `Method not allowed`
	if rr.Body.String() != want {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), want)
	}
}

func TestFibHandlerRetrieveErrorMethodDeleteNotAllowed(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/fib?n=72", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	fibApiImpl := NewFibApiImpl()
	handler := http.HandlerFunc(fibApiImpl.fibonacciResourcesHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	want := `Method not allowed`
	if rr.Body.String() != want {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), want)
	}
}

func TestFibHandlerRetrieveErrorMethodPutNotAllowed(t *testing.T) {
	req, err := http.NewRequest("PUT", "/fib?n=72", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	fibApiImpl := NewFibApiImpl()
	handler := http.HandlerFunc(fibApiImpl.fibonacciResourcesHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	want := `Method not allowed`
	if rr.Body.String() != want {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), want)
	}
}

func TestFibHandlerRetrieveErrorMethodPostNotAllowed(t *testing.T) {
	req, err := http.NewRequest("POST", "/fib?n=72", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	fibApiImpl := NewFibApiImpl()
	handler := http.HandlerFunc(fibApiImpl.fibonacciResourcesHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	want := `Method not allowed`
	if rr.Body.String() != want {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), want)
	}
}

func TestFibHandlerRetrieveError(t *testing.T) {
	req, err := http.NewRequest("GET", "/fib", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	fibApiImpl := NewFibApiImpl()
	handler := http.HandlerFunc(fibApiImpl.fibResourcesHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	want := "ERROR: required query parameted 'N' was not provided"
	got := rr.Body.String()
	if got != want {
		t.Errorf("handler returned unexpected body: got %v want %v",
			got, want)
	}
}

func TestFibHandlerRetrieveInvalidParameter(t *testing.T) {
	req, err := http.NewRequest("GET", "/fib?n=-1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	fibApiImpl := NewFibApiImpl()
	handler := http.HandlerFunc(fibApiImpl.fibResourcesHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	want := "ERROR: invalid parameter range. Should be between 0 - 1474, but was: -1"
	got := rr.Body.String()
	if got != want {
		t.Errorf("handler returned unexpected body: got %v want %v",
			got, want)
	}
}

func TestFibHandlerRetrieveInvalidParameter1475(t *testing.T) {
	req, err := http.NewRequest("GET", "/fib?n=1475", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	fibApiImpl := NewFibApiImpl()
	handler := http.HandlerFunc(fibApiImpl.fibResourcesHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	want := "ERROR: invalid parameter range. Should be between 0 - 1474, but was: 1475"
	got := rr.Body.String()
	if got != want {
		t.Errorf("handler returned unexpected body: got %v want %v",
			got, want)
	}
}

func TestFibHandlerRetrieveInvalidParameterFoo(t *testing.T) {
	req, err := http.NewRequest("GET", "/fib?n=Foo", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	fibApiImpl := NewFibApiImpl()
	handler := http.HandlerFunc(fibApiImpl.fibResourcesHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	want := "ERROR: invalid input value Foo"
	got := rr.Body.String()
	if got != want {
		t.Errorf("handler returned unexpected body: got %v want %v",
			got, want)
	}
}

func TestFibHandlerRetrieveError2(t *testing.T) {
	req, err := http.NewRequest("GET", "/fib?m=santana", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	fibApiImpl := NewFibApiImpl()
	handler := http.HandlerFunc(fibApiImpl.fibResourcesHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	want := "ERROR: required query parameted 'N' was not provided"
	got := rr.Body.String()
	if got != want {
		t.Errorf("handler returned unexpected body: got %v want %v",
			got, want)
	}
}

func TestFibHandlerRetrieve72(t *testing.T) {
	req, err := http.NewRequest("GET", "/fib?n=72", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	fibApiImpl := NewFibApiImpl()
	handler := http.HandlerFunc(fibApiImpl.fibResourcesHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	want := "498454011879264"
	got := rr.Body.String()
	if got != want {
		t.Errorf("handler returned unexpected body: got %v want %v",
			got, want)
	}
}
