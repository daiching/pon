package pon

import (
	"fmt"
	"net/http"
)

// It is wanted to set default json without particular statuscode specified by user by status code.
// So, default status code(normal:StatusNormal, error:StatusUnnormal) is prepared in this package.
// Also, general status code (int) is prepared having compatibility with http package(except for some).
const (
	// status codes (int)
	// If next two status is selected by user, default status code is set according to method both normal and error.
	StatusNormal   int = 0
	StatusUnnormal int = -1

	// 2XX (now, 200~208)
	StatusOK                   int = http.StatusOK
	StatusCreated              int = http.StatusCreated
	StatusAccepted             int = http.StatusAccepted
	StatusNonAuthoritativeInfo int = http.StatusNonAuthoritativeInfo
	StatusNoContent            int = http.StatusNoContent
	StatusResetContent         int = http.StatusResetContent
	StatusPartialContent       int = http.StatusPartialContent
	StatusMultiStatus          int = http.StatusMultiStatus
	StatusAlreadyReported      int = http.StatusAlreadyReported

	// 4XX(now, 400~417)
	StatusBadRequest                   int = http.StatusBadRequest
	StatusUnauthorized                 int = http.StatusUnauthorized
	StatusPaymentRequired              int = http.StatusPaymentRequired
	StatusForbidden                    int = http.StatusForbidden
	StatusNotFound                     int = http.StatusNotFound
	StatusMethodNotAllowed             int = http.StatusMethodNotAllowed
	StatusNotAcceptable                int = http.StatusNotAcceptable
	StatusProxyAuthRequired            int = http.StatusProxyAuthRequired
	StatusRequestTimeout               int = http.StatusRequestTimeout
	StatusConflict                     int = http.StatusConflict
	StatusGone                         int = http.StatusGone
	StatusLengthRequired               int = http.StatusLengthRequired
	StatusPreconditionFailed           int = http.StatusPreconditionFailed
	StatusRequestEntityTooLarge        int = http.StatusRequestEntityTooLarge
	StatusRequestURITooLong            int = http.StatusRequestURITooLong
	StatusUnsupportedMediaType         int = http.StatusUnsupportedMediaType
	StatusRequestedRangeNotSatisfiable int = http.StatusRequestedRangeNotSatisfiable
	StatusExpectationFailed            int = http.StatusExpectationFailed

	// 5XX(now, 500~505)
	StatusInternalServerError     int = http.StatusInternalServerError
	StatusNotImplemented          int = http.StatusNotImplemented
	StatusBadGateway              int = http.StatusBadGateway
	StatusServiceUnavailable      int = http.StatusServiceUnavailable
	StatusGatewayTimeout          int = http.StatusGatewayTimeout
	StatusHTTPVersionNotSupported int = http.StatusHTTPVersionNotSupported
)

// 1. only http status 2XX, 4XX or 5XX is allowed.
// 2. if json is blank and http status is error, default error json is returned.
func permeatedDefaultFilter(js string, sts int) (string, int) {
	if sts >= 300 && sts < 400 {
		fmt.Println("Reesponse is 3XX.")
		return getDefaultErrorJson(StatusInternalServerError), StatusInternalServerError
	}
	if js == "" && sts >= 400 {
		if sts >= 500 {
			fmt.Println("Reesponse is 5XX.")
		} else {
			fmt.Println("Reesponse is 4XX.")
		}
		return getDefaultErrorJson(sts), sts
	}
	return js, sts
}

// if StatusNormal or StatusUnnormal is set, default HTTP status code was returned depending on method
func getDefaultStatusWithMethod(m string, sts int) int {
	if !(sts == StatusNormal || sts == StatusUnnormal) {
		return sts
	}

	switch m {
	case "GET":
		if sts == StatusNormal {
			return StatusOK
		} else {
			return StatusBadRequest
		}
	case "POST":
		if sts == StatusNormal {
			return StatusCreated
		} else {
			return StatusConflict
		}
	case "PUT":
		if sts == StatusNormal {
			return StatusNoContent
		} else {
			return StatusConflict
		}
	case "DELETE":
		if sts == StatusNormal {
			return StatusNoContent
		} else {
			return StatusConflict
		}
	default:
		return StatusMethodNotAllowed
	}
}
