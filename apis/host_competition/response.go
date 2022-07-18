package validate_phone_number_api

import (
	"encoding/json"
	"find_competitor/common"
	"net/http"
)

type BaseSuccessResponse struct {
	IsValidated bool `json:"is_validated"`
}

type SuccessResponse struct {
	Data BaseSuccessResponse `json:"data"`
	common.BaseResponse
}

/* Response Functions */
func generateSuccessResponse(requestUrl string, isValidated bool, w http.ResponseWriter) {
	successResponse := SuccessResponse{
		Data:         BaseSuccessResponse{IsValidated: isValidated},
		BaseResponse: common.BaseResponse{StatusCode: 200, Success: true, Cmd: requestUrl},
	}
	json.NewEncoder(w).Encode(&successResponse)
}
