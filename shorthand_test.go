package simplehttp

import (
	"os"
	"testing"
)

func reqEnv(key string, cb func(string)) {
	value := os.Getenv(key)
	if value != "" && cb != nil {
		cb(value)
	}
}

func TestSimpleGet(t *testing.T) {
	reqEnv("SIMPLEHTTP_TEST_SERVER", func(server string) {
		Request{
			Url: "http://" + server + "/get",
		}.Get(func(data []byte) {
			t.Log(string(data))
		}, func(data []byte) {
			t.Log(string(data))
		})
	})
}

func TestSimplePost(t *testing.T) {
	reqEnv("SIMPLEHTTP_TEST_SERVER", func(server string) {
		Request{
			Url:  "http://" + server + "/post",
			Data: []byte("some data"),
		}.Post(func(data []byte) {
			t.Log(string(data))
		}, func(data []byte) {
			t.Log(string(data))
		})
	})
}
