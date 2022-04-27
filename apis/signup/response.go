package signup_api

import "find_competitor/common"

type SuccessResponse struct {
	IsSignedUp bool `json:"is_signed_up"`
	common.BaseResponse
}
