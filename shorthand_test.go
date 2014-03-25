package simplehttp

import (
	"testing"
)

const (
	simpleGetReturn = "Welcome to the real world!"
)

func TestSimpleGet(t *testing.T) {
	Request{
		Url:    "http://localhost:8888/",
		Method: "GET",
	}.Do()
}
