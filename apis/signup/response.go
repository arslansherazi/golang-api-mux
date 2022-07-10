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

/* Response Functions */
func generateSignupSuccessResponse(requestUrl string, w http.ResponseWriter) {
	successResponse := SuccessResponse{
		IsSignedUp:   true,
		BaseResponse: common.BaseResponse{StatusCode: 200, Success: true, Cmd: requestUrl},
	}
	json.NewEncoder(w).Encode(&successResponse)
}
