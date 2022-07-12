package login_api

import (
	"encoding/json"
	"find_competitor/common"
	"net/http"
)

type SuccessResponse struct {
	IsLoggedIn      bool   `json:"is_loggedin"`
	Name            string `json:"name"`
	ProfileImageUrl string `json:"profile_image_url"`
	Token           string `json:"token"`
	common.BaseResponse
}

/* Response Functions */
func generateSuccessResponse(requestUrl string, name string, profileImageUrl string, token string, w http.ResponseWriter) {
	successResponse := SuccessResponse{
		IsLoggedIn:      true,
		Name:            name,
		ProfileImageUrl: profileImageUrl,
		Token:           token,
		BaseResponse:    common.BaseResponse{StatusCode: 200, Success: true, Cmd: requestUrl},
	}
	json.NewEncoder(w).Encode(&successResponse)
}
