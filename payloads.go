package simplehttp

type FormDataPayload struct {
	Values map[string]string
	Files  map[string]string
}

type UrlEncodedPayload struct {
	Values map[string]string
}

type RawPayload struct {
	Data []byte
}
