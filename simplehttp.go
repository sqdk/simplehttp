package simplehttp

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

type SimpleHTTPRequest struct {
	Method            string
	URL               string
	Parameters        map[string]string
	FormValues        map[string]string
	Headers           map[string]string
	BasicAuthUser     string
	BasicAuthPassword string
}

func NewSimpleHTTPRequest(method, url string) *SimpleHTTPRequest {
	return &SimpleHTTPRequest{Method: method, URL: url}
}

func NewGetRequest(url string) *SimpleHTTPRequest {
	return NewSimpleHTTPRequest("GET", url)
}

func NewPostRequest(url string) *SimpleHTTPRequest {
	return NewSimpleHTTPRequest("POST", url)
}

func NewDeleteRequest(url string) *SimpleHTTPRequest {
	return NewSimpleHTTPRequest("DELETE", url)
}

func (r *SimpleHTTPRequest) AddParameter(name, value string) {
	if r.Parameters == nil {
		r.Parameters = make(map[string]string)
	}
	r.Parameters[name] = value
}

func (r *SimpleHTTPRequest) AddFormValue(name, value string) {
	if r.FormValues == nil {
		r.FormValues = make(map[string]string)
	}
	r.FormValues[name] = value
}

func (r *SimpleHTTPRequest) AddHeader(name, value string) {
	if r.Headers == nil {
		r.Headers = make(map[string]string)
	}
	r.Headers[name] = value
}

func (r *SimpleHTTPRequest) SetBasicAuth(user, password string) {
	r.BasicAuthUser = user
	r.BasicAuthPassword = password
}

func (r *SimpleHTTPRequest) MakeJSONRequest(v interface{}) error {
	responseBody, err := r.MakeRequest()
	if err != nil {
		return err
	}
	return json.Unmarshal(responseBody, v)
}

func (r *SimpleHTTPRequest) MakeRequest() ([]byte, error) {
	url, err := r.generateUrlWithParameters()
	if err != nil {
		return make([]byte, 0), err
	}
	bodyData, hasBody := r.makeBodyData()

	var body io.Reader
	if hasBody {
		body = bytes.NewBufferString(bodyData.Encode())
	} else {
		body = nil
	}

	req, err := http.NewRequest(r.Method, url, body)
	if err != nil {
		return make([]byte, 0), err
	}

	if r.BasicAuthUser != "" && r.BasicAuthPassword != "" {
		req.SetBasicAuth(r.BasicAuthUser, r.BasicAuthPassword)
	}

	for header, value := range r.Headers {
		req.Header.Add(header, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return make([]byte, 0), err
	}

	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return make([]byte, 0), err
	}

	return responseBody, nil
}

func (r *SimpleHTTPRequest) makeBodyData() (data url.Values, hasBody bool) {
	data = url.Values{}
	if r.FormValues != nil && len(r.FormValues) > 0 {
		hasBody = true
		r.AddHeader("Content-Type", "application/x-www-form-urlencoded")
		for name, value := range r.FormValues {
			data.Add(name, value)
		}
	} else {
		hasBody = false
	}

	return
}

func (r *SimpleHTTPRequest) generateUrlWithParameters() (string, error) {
	url, err := url.Parse(r.URL)
	if err != nil {
		return "", err
	}
	q := url.Query()
	if r.Parameters != nil && len(r.Parameters) > 0 {
		for name, value := range r.Parameters {
			q.Set(name, value)
		}
	}
	url.RawQuery = q.Encode()

	return url.String(), nil
}
