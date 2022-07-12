package login_api

import (
	"encoding/json"
	"find_competitor/common"
	"net/http"
)

type BaseSuccessResponse struct {
	IsLoggedIn      bool   `json:"is_logged_in"`
	UserID          int    `json:user_id`
	Name            string `json:"name"`
	ProfileImageUrl string `json:"profile_image_url"`
	Token           string `json:"token"`
}

type SuccessResponse struct {
	Data BaseSuccessResponse `json:"data"`
	common.BaseResponse
}

/* Response Functions */
func generateSuccessResponse(requestUrl string, userID int, name string, profileImageUrl string, token string, w http.ResponseWriter) {
	successResponse := SuccessResponse{
		Data: BaseSuccessResponse{
			IsLoggedIn:      true,
			UserID:          userID,
			Name:            name,
			ProfileImageUrl: profileImageUrl,
			Token:           token,
		},
		BaseResponse: common.BaseResponse{StatusCode: 200, Success: true, Cmd: requestUrl},
	}
	json.NewEncoder(w).Encode(&successResponse)
}
