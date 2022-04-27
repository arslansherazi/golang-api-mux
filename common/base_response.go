package common

type BaseResponse struct {
	StatusCode uint32 `json:"status_code"`
	Success    bool   `json:"success"`
	Cmd        string `json:"cmd"`
}
