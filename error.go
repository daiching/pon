package pon

import (
	"encoding/json"
	"net/http"
)

// This is default struct in error.
// If user specify any error data, json Marshaled from this sturct is returned in error.
type ErrorResponse struct {
	status       int    `json:"status"`
	errorMessage string `json:"errorMessage"`
}

func NewErrorResponse(err string, sts int) ErrorResponse {
	res := ErrorResponse{
		StatusBadRequest,
		http.StatusText(StatusBadRequest),
	}
	if err != "" && sts != 0 {
		res.errorMessage = err
		return res
	}
	if err == "" && sts != 0 {
		res.status = sts
		res.errorMessage = http.StatusText(sts)
		return res
	}
	res.status = sts
	res.errorMessage = err
	return res
}

func NewErrorResponseJson(err string, sts int) string {
	b, _ := json.Marshal(NewErrorResponse(err, sts))
	return string(b)
}

func NewErrorResponseAndStatus(err string, sts int) (ErrorResponse, int) {
	res := NewErrorResponse(err, sts)
	if sts == 0 {
		return res, StatusBadRequest
	}
	return res, sts
}

func NewErrorResponseJsonAndStatus(err string, sts int) (string, int) {
	res, sts := NewErrorResponseAndStatus(err, sts)
	b, _ := json.Marshal(res)
	return string(b), sts
}

func (e *ErrorResponse) SetStatus(sts int) {
	e.status = sts
}

func (e *ErrorResponse) SetErrorMessage(em string) {
	e.errorMessage = em
}

func getDefaultErrorJson(sts int) string {
	e := ErrorResponse{
		sts,
		http.StatusText(sts),
	}
	b, _ := json.Marshal(e)
	return string(b)
}
