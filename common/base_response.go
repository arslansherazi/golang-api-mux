package common

import (
	"encoding/json"
	"net/http"
)

type BaseResponse struct {
	StatusCode uint32 `json:"status_code"`
	Success    bool   `json:"success"`
	Cmd        string `json:"cmd"`
}

/* Base response functions */
type ErrorResponse struct {
	Message string `json:"message"`
	BaseResponse
}

func Generate500ErrorResponse(requestUrl string, w http.ResponseWriter) {
	errorResponse := ErrorResponse{
		Message:      "Internal Server Error",
		BaseResponse: BaseResponse{StatusCode: 500, Success: false, Cmd: requestUrl},
	}
	json.NewEncoder(w).Encode(&errorResponse)
}

func Generate422Response(requestUrl string, message string, w http.ResponseWriter) {
	errorResponse := ErrorResponse{
		Message:      message,
		BaseResponse: BaseResponse{StatusCode: 422, Success: false, Cmd: requestUrl},
	}
	json.NewEncoder(w).Encode(&errorResponse)
}
