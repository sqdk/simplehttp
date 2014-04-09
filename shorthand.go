package simplehttp

type BasicAuthentication struct {
	User     string
	Password string
}

// The function prototype that handles responses using the
// shorthand notation.
type shorthandResponseHandler func([]byte)

type Request struct {
	Url            string
	Authentication BasicAuthentication
	UserAgent      string
}

func createHttpRequest(req Request) *HTTPRequest {
	r := NewHTTPRequest(req.Url)
	if req.Authentication.User != "" {
		r.SetBasicAuth(req.Authentication.User, req.Authentication.Password)
	}
	if req.UserAgent != "" {
		r.AddHeader("User-Agent", req.UserAgent)
	}
	return r
}

func (r Request) Get(success, failure shorthandResponseHandler) {
	req := createHttpRequest(r)
	res, err := req.MakeGetRequest()

	if err != nil || res.Code >= 400 {
		if failure != nil {
			failure(res.Data)
		}
	} else {
		if success != nil {
			success(res.Data)
		}
	}
}

func (r Request) Post(success, failure shorthandResponseHandler) {
	req := createHttpRequest(r)
	res, err := req.MakePostRequest(nil)

	if err != nil || res.Code >= 400 {
		if failure != nil {
			failure(res.Data)
		}
	} else {
		if success != nil {
			success(res.Data)
		}
	}
}
