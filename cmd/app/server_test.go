package main

import (
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"testing"
	"time"
)

func TestServer_Root(t *testing.T) {
	ts := testServer()
	defer ts.Close()

	path := "/"
	res, err := ts.Client().Get(fmt.Sprintf("%s%s", ts.URL, path))
	if err != nil {
		t.Fatalf("Failed to GET '%s': %v", path, err)
	}

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Fatalf("Failed to read body: %v", err)
	}

	got := string(body)
	want := "Hello World\n"
	if got != want {
		t.Fatalf("Unexpected response body:\nwant:\n%s\ngot:\n%s\n", want, got)
	}
}

func TestServer_Content(t *testing.T) {
	ts := testServer()
	defer ts.Close()

	path := "/content.json"
	res, err := ts.Client().Get(fmt.Sprintf("%s%s", ts.URL, path))
	if err != nil {
		t.Fatalf("Failed to GET '%s': %v", path, err)
	}

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Fatalf("Failed to read body: %v", err)
	}

	got := string(body)
	want := "{\"Key\":\"value\"}\n"
	if got != want {
		t.Fatalf("Unexpected response body:\nwant:\n%s\ngot:\n%s\n", want, got)
	}
}

func TestServer_Healthz(t *testing.T) {
	ts := testServer()
	defer ts.Close()

	path := "/healthz"
	res, err := ts.Client().Get(fmt.Sprintf("%s%s", ts.URL, path))
	if err != nil {
		t.Fatalf("Failed to GET '%s': %v", path, err)
	}

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Fatalf("Failed to read body: %v", err)
	}

	got := string(body)
	want := "healthy"
	if got != want {
		t.Fatalf("Unexpected response body:\nwant:\n%s\ngot:\n%s\n", want, got)
	}
}

func testServer() *httptest.Server {
	return httptest.NewServer(NewServer("", 0, 0*time.Second).Handler())
}
