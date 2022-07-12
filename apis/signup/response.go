package signup_api

import (
	"encoding/json"
	"find_competitor/common"
	"net/http"
)

type BaseSuccessResponse struct {
	IsSignedUp bool `json:"is_signed_up"`
}

type SuccessResponse struct {
	Data BaseSuccessResponse `json:"data"`
	common.BaseResponse
}

/* Response Functions */
func generateSuccessResponse(requestUrl string, w http.ResponseWriter) {
	successResponse := SuccessResponse{
		Data: BaseSuccessResponse{
			IsSignedUp: true,
		},
		BaseResponse: common.BaseResponse{StatusCode: 200, Success: true, Cmd: requestUrl},
	}
	json.NewEncoder(w).Encode(&successResponse)
}
