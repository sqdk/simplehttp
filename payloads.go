package simplehttp

import (
	"errors"
)

type Payload interface {
	GetPayloadBytes() []byte
}

type FormDataPayload struct {
	UrlEncodedPayload
	Files map[string]string
}

type UrlEncodedPayload struct {
	Values map[string]string
}

type RawPayload struct {
	Data []byte
}

func NewFormDataPayload() *FormDataPayload {
	return &FormDataPayload{}
}

func (f *FormDataPayload) AddFile(key, file string) error {
	if _, ok := f.Files[key]; !ok {
		f.Files[key] = file
		return nil
	} else {
		return errors.New("File already exists.")
	}
}

func (f *FormDataPayload) RemoveFile(key string) error {
	if _, ok := f.Files[key]; !ok {
		delete(f.Files, key)
		return nil
	} else {
		return errors.New("File doesn't exist.")
	}
}

func (f *FormDataPayload) GetPayloadBytes() []byte {
	return nil
}

func NewUrlEncodedPayload() *UrlEncodedPayload {
	return &UrlEncodedPayload{}
}

func (f *UrlEncodedPayload) AddValue(key, value string) error {
	if _, ok := f.Values[key]; !ok {
		f.Values[key] = value
		return nil
	} else {
		return errors.New("Value already exists.")
	}
}

func (f *UrlEncodedPayload) RemoveValue(key string) error {
	if _, ok := f.Values[key]; !ok {
		delete(f.Values, key)
		return nil
	} else {
		return errors.New("Value doesn't exist.")
	}
}

func (f *UrlEncodedPayload) GetPayloadBytes() []byte {
	return nil
}
