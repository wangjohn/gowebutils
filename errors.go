package gowebutils

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/wangjohn/gowebutils/jsonapi"
)

func SendError(w http.ResponseWriter, err error) {
	e := jsonapi.Error{
		Status: "400",
		Code:   "error",
		Detail: err.Error(),
	}
	errorList := []jsonapi.Error{e}
	resp := jsonapi.ErrorResponse{
		Errors: errorList,
	}
	log.Printf("Returning error response: %+v", resp)
	json.NewEncoder(w).Encode(resp)
}
