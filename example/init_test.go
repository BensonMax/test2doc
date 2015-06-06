package main

import (
	"net/http/httptest"
	"testing"

	"github.com/adams-sarah/prettytest"
	"github.com/adams-sarah/test2doc/test"
)

var server *httptest.Server

type mainSuite struct {
	prettytest.Suite
}

func TestRunner(t *testing.T) {
	var err error

	server, err = test.NewServer(newMux(), ".")
	if err != nil {
		panic(err.Error())
	}
	defer server.Close()

	prettytest.RunWithFormatter(
		t,
		new(prettytest.TDDFormatter),
		new(mainSuite),
	)
}
