package gowebutils

import (
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
