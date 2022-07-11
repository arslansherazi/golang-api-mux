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
type BaseErrorResponse struct {
	Message string `json:"message"`
	BaseResponse
}

func ErrorResponse(requestUrl string, statusCode uint32, message string, w http.ResponseWriter) {
	errorResponse := BaseErrorResponse{
		Message:      message,
		BaseResponse: BaseResponse{StatusCode: statusCode, Success: false, Cmd: requestUrl},
	}
	json.NewEncoder(w).Encode(&errorResponse)
}
