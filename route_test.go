package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(homeHandler))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Fatalf("Unexpected StatusCode: %d\n", res.StatusCode)
	}

	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
}
