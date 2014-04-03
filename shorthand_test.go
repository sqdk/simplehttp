package simplehttp

import (
	"testing"
)

func TestSimpleGet(t *testing.T) {
	Request{
		Url: "http://localhost:4000/foobar",
	}.Get(func(data []byte) {
		t.Log(string(data))
	}, func(data []byte) {
		t.Log(string(data))
	})
}
