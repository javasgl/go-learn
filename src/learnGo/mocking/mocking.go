package mocking

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

var msg = `this is msg`

func MockServer() *httptest.Server {

	hanlder := func(w http.ResponseWriter, r *http.Request) {

		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/xml")
		fmt.Fprintln(w, msg)
	}
	return httptest.NewServer(http.HandlerFunc(hanlder))
}

func TestServer(t *testing.T) {

	status := http.StatusOK
	server := MockServer()
	defer server.Close()

	t.Log("start to test mock server", server.URL)
	resp, err := http.Get(server.URL)
	if err != nil {
		t.Fatal("error,cant`t access ", server.URL)
	}
	defer resp.Body.Close()
	if resp.StatusCode != status {
		t.Fatal("stuatus !=200", resp.StatusCode)
	}
}

func ExampleMockServer() {

	fmt.Println("test")
	// Output:
	// test
}
