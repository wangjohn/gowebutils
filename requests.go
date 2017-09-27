package gowebutils

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

const MaxRequestSize = 1048576 * 5

func PrepareRequestBody(r *http.Request) ([]byte, error) {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, MaxRequestSize))
	if err != nil {
		return body, err
	}
	err = r.Body.Close()
	return body, err
}

func UnmarshalRequestBody(r *http.Request, v interface{}) error {
	body, err := PrepareRequestBody(r)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(body, v); err != nil {
		return err
	}
	return nil
}
