package main

import (
	"fmt"
	"net/http"
	"testing"

	"greenlight.karronqiu.github.com/internal/data"
	"greenlight.karronqiu.github.com/internal/data/mocks"
)

func TestCreateMovieHandler(t *testing.T) {
	app := newTestApplication(t)

	ts := newTestServer(t, app.routes())
	defer ts.Close()

	h := make(map[string]string)
	h["Authorization"] = "Bearer Y3QMGX3PJ3WLRL2YRTQGQ6KRHU"

	input := struct {
		Title   string       `json:"title"`
		Year    int32        `json:"year"`
		Runtime data.Runtime `json:"runtime"`
		Genres  []string     `json:"genres"`
	}{
		Title:   mocks.MockMovie.Title,
		Year:    mocks.MockMovie.Year,
		Genres:  mocks.MockMovie.Genres,
		Runtime: mocks.MockMovie.Runtime,
	}

	s, header, body := ts.post(t, "/v1/movies/", input, h)

	assertEquals(t, http.StatusCreated, s, "status code: ")
	assertEquals(t, fmt.Sprintf("/v1/movies/%d", mocks.MockMovie.ID), header.Get("Location"), "Location: ")
	assertTrue(t, len(body) > 0, "body: ")
}
