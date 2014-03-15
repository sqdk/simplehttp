package simplehttp

import (
	"bytes"
	"errors"
	"io"
	"mime/multipart"
	"net/url"
	"os"
	"path"
)

type Payload interface {
	GetPayloadBuffer() (*bytes.Buffer, error)
	GetContentType() string
}

type FormDataPayload struct {
	contentType string
	Values      map[string]string
	Files       map[string]string
}

type UrlEncodedPayload struct {
	Values map[string]string
}

type RawPayload struct {
	Data []byte
}

func NewFormDataPayload() *FormDataPayload {
	values := make(map[string]string)
	files := make(map[string]string)
	return &FormDataPayload{Values: values, Files: files}
}

func (f *FormDataPayload) AddValue(key, value string) error {
	if _, ok := f.Values[key]; !ok {
		f.Values[key] = value
		return nil
	} else {
		return errors.New("Value already exists.")
	}
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

func (f *FormDataPayload) GetPayloadBuffer() (*bytes.Buffer, error) {
	data := &bytes.Buffer{}
	writer := multipart.NewWriter(data)
	defer writer.Close()

	for name, value := range f.Values {
		if tmp, err := writer.CreateFormField(name); err == nil {
			tmp.Write([]byte(value))
		} else {
			return nil, err
		}
	}

	for name, file := range f.Files {
		if tmp, err := writer.CreateFormFile(name, path.Base(file)); err == nil {
			if fp, err := os.Open(file); err == nil {
				defer fp.Close()
				io.Copy(tmp, fp)
			} else {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	f.contentType = writer.FormDataContentType()

	return data, nil
}

func (f *FormDataPayload) GetContentType() string {
	if f.contentType == "" {
		f.GetPayloadBuffer()
	}
	return f.contentType
}

func NewUrlEncodedPayload() *UrlEncodedPayload {
	values := make(map[string]string)
	return &UrlEncodedPayload{Values: values}
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

func (f *UrlEncodedPayload) GetPayloadBuffer() (*bytes.Buffer, error) {
	data := url.Values{}
	for key, value := range f.Values {
		data.Add(key, value)
	}
	return bytes.NewBufferString(data.Encode()), nil
}

func (f *UrlEncodedPayload) GetContentType() string {
	return "application/x-www-form-urlencoded"
}
