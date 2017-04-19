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
	sendErrorList(w, errorList)
}

func SendErrors(w http.ResponseWriter, errors []error) {
	errorList := []jsonapi.Error{}
	for _, e := range errors {
		errorList = append(errorList, jsonapi.Error{
			Status: "400",
			Code:   "error",
			Detail: e.Error(),
		})
	}
	sendErrorList(w, errorList)
}

func SendErrorMessages(w http.ResponseWriter, errorMessages []string) {
	errorList := []jsonapi.Error{}
	for _, msg := range errorMessages {
		errorList = append(errorList, jsonapi.Error{
			Status: "400",
			Code:   "error",
			Detail: msg,
		})
	}
	sendErrorList(w, errorList)
}

func sendErrorList(w http.ResponseWriter, errorList []jsonapi.Error) {
	resp := jsonapi.ErrorResponse{Errors: errorList}
	log.Printf("Returning error response: %+v", resp)
	json.NewEncoder(w).Encode(resp)
}
