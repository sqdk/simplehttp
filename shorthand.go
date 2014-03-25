package simplehttp

import "fmt"

type BasicAuthentication struct {
	User     string
	Password string
}

type Request struct {
	Url            string
	Method         string
	Authentication BasicAuthentication
}

func (r Request) Do() {
	fmt.Printf("%#v\n", r.Authentication)

	if r.Method == "GET" {
		doGet(r)
	}
}

func doGet(r Request) {
	req := NewHTTPRequest(r.Url)
	req.MakeGetRequest()
}
