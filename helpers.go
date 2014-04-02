package simplehttp

func (r *HTTPRequest) GetResponseFromJSON(v interface{}) (int, error) {
	response, err := r.MakeGetRequest()
	if err != nil {
		return response.Code, err
	}
	return response.Code, response.ParseFromJSON(v)
}

func (r *HTTPRequest) PostResponseFromJSON(payload Payload, v interface{}) (int, error) {
	response, err := r.MakePostRequest(payload)
	if err != nil {
		return response.Code, err
	}
	return response.Code, response.ParseFromJSON(v)
}
