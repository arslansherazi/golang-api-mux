package signup_api

import (
	"encoding/json"
	"find_competitor/common"
	"net/http"
)

type SuccessResponse struct {
	IsSignedUp bool `json:"is_signed_up"`
	common.BaseResponse
}

type ErrorResponse struct {
	Message string `json:"message"`
	common.BaseResponse
}

/* Response Functions */
func generateSuccessResponse(requestUrl string, w http.ResponseWriter) {
	successResponse := SuccessResponse{
		IsSignedUp:   true,
		BaseResponse: common.BaseResponse{StatusCode: 200, Success: true, Cmd: requestUrl},
	}
	json.NewEncoder(w).Encode(&successResponse)
}

func generate500ErrorResponse(requestUrl string, w http.ResponseWriter) {
	errorResponse := ErrorResponse{
		Message:      "Internal Server Error",
		BaseResponse: common.BaseResponse{StatusCode: 500, Success: false, Cmd: requestUrl},
	}
	json.NewEncoder(w).Encode(&errorResponse)
}

func generate422Response(requestUrl string, message string, w http.ResponseWriter) {
	errorResponse := ErrorResponse{
		Message:      message,
		BaseResponse: common.BaseResponse{StatusCode: 422, Success: false, Cmd: requestUrl},
	}
	json.NewEncoder(w).Encode(&errorResponse)
}
