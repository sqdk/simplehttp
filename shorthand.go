package simplehttp

// Type to encapsulate basic authentication for requests.
type BasicAuthentication struct {
	User     string
	Password string
}

// The function prototype that handles responses using the
// shorthand notation.
type shorthandResponseHandler func([]byte)

// Type to wrap requests.
type Request struct {
	Url            string
	Authentication BasicAuthentication
	UserAgent      string
	Data           []byte
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

// Function that performs a basic GET request and calls the proper
// callback functions.
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

// Function that performs a POST request and calls the proper
// callback functions.
func (r Request) Post(success, failure shorthandResponseHandler) {
	req := createHttpRequest(r)
	var payload Payload = nil
	if r.Data != nil {
		payload = NewRawPayload(r.Data)
	}
	res, err := req.MakePostRequest(payload)

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
