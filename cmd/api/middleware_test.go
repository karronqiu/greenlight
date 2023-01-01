package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEnableCORS(t *testing.T) {
	rr := httptest.NewRecorder()

	// Mock a preflight request
	r, err := http.NewRequest(http.MethodOptions, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	app := &application{}
	origin := "http://localhost:4000"
	app.config.cors.trustedOrigins = []string{origin}

	r.Header.Add("Origin", origin)
	r.Header.Add("Access-Control-Request-Method", "PUT") // request a PUT method

	// Create a mock HTTP handler
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	app.enableCORS(next).ServeHTTP(rr, r)

	rs := rr.Result()
	actual := rs.Header.Values("Vary")

	assertEquals(t, "Origin", actual[0], "header: Origin ")
	assertEquals(t, "Access-Control-Request-Method", actual[1], "header: Access-Control-Request-Method")

	assertEquals(t, origin, rs.Header.Get("Access-Control-Allow-Origin"), "header: Access-Control-Allow-Origin")
	assertEquals(t, "OPTIONS, PUT, PATCH, DELETE", rs.Header.Get("Access-Control-Allow-Methods"), "header: Access-Control-Allow-Methods")
	assertEquals(t, "Authorization, Content-Type", rs.Header.Get("Access-Control-Allow-Headers"), "header: Access-Control-Allow-Headers")

	assertEquals(t, http.StatusOK, rs.StatusCode, "status code: ")

	defer rs.Body.Close()
	body, err := io.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}

	// The response body is empty returned by the middleware.
	// It is not returned by the mock http handler because the middleware responded the OPTION method directly.
	assertEquals(t, 0, len(body), "body length: ")
}
