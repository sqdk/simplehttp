package simplehttp

import (
	"encoding/json"
	"encoding/xml"
)

func (r *HTTPResponse) ParseFromJSON(v interface{}) error {
	return json.Unmarshal(r.Data, v)
}

func (r *HTTPResponse) ParseFromJSONToFirst(v []interface{}) int {
	for i, val := range v {
		if err := json.Unmarshal(r.Data, val); err == nil {
			return i
		}
	}
	return -1
}

func (r *HTTPResponse) ParseFromXML(v interface{}) error {
	return xml.Unmarshal(r.Data, v)
}

func (r *HTTPResponse) ParseToXMToFirst(v []interface{}) int {
	for i, val := range v {
		if err := xml.Unmarshal(r.Data, val); err == nil {
			return i
		}
	}
	return -1
}
