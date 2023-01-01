package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"greenlight.karronqiu.github.com/internal/data/mocks"
	"greenlight.karronqiu.github.com/internal/jsonlog"
)

type testServer struct {
	*httptest.Server
}

func newTestServer(t *testing.T, h http.Handler) *testServer {
	ts := httptest.NewServer(h)
	return &testServer{ts}
}

func newTestApplication(t *testing.T) *application {
	var cfg config
	cfg.cors.trustedOrigins = []string{"http://localhost:4000"}
	cfg.env = "development"
	cfg.db.dsn = "postgres://greenlight:hustpower@192.168.50.4:15432?sslmode=disable"
	cfg.db.maxOpenConns = 25
	cfg.db.maxIdleConns = 25
	cfg.db.maxIdleTime = "15m"
	cfg.limiter.rps = 4
	cfg.limiter.burst = 4
	cfg.limiter.enabled = true
	logger := jsonlog.New(os.Stdout, jsonlog.LevelInfo)

	return &application{
		config: cfg,
		logger: logger,
		models: mocks.NewModels(),
	}
}

//lint:ignore U1000 Ignore unused function temporarily for debugging
func (ts *testServer) get(t *testing.T, urlPath string) (int, http.Header, string) {
	rs, err := ts.Client().Get(ts.URL + urlPath)
	if err != nil {
		t.Fatal(err)
	}

	defer rs.Body.Close()

	body, err := io.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}

	bytes.TrimSpace(body)

	return rs.StatusCode, rs.Header, string(body)
}

func (ts *testServer) post(t *testing.T, urlPath string, body any, headers map[string]string) (int, http.Header, string) {
	b, err := json.Marshal(body)
	if err != nil {
		t.Fatal(err)
	}

	r, err := http.NewRequest(http.MethodPost, ts.URL+urlPath, bytes.NewReader(b))
	if err != nil {
		t.Fatal(err)
	}

	r.Header.Add("Content-Type", "application/json")
	for key, value := range headers {
		r.Header.Add(key, value)
	}

	rs, err := ts.Client().Do(r)
	if err != nil {
		t.Error(err)
	}

	defer rs.Body.Close()

	rsBody, err := io.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}

	bytes.TrimSpace(rsBody)

	return rs.StatusCode, rs.Header, string(rsBody)
}

func assertEquals[T comparable](t *testing.T, expectedValue, actualValue T, message string) {
	t.Helper()

	if actualValue != expectedValue {
		t.Errorf("%s, got: %v; want: %v", message, actualValue, expectedValue)
	}
}

func assertTrue(t *testing.T, r bool, message string) {
	t.Helper()

	if !r {
		t.Errorf("%s, Expected true, actual false", message)
	}
}
