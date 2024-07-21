package types

type ApiResponse struct {
	ResultCode  int64       `json:"resultCode"` // 1 : success / -1 : fail
	Description string      `json:"description"`
	ErrorCode   interface{} `json:"errCode"`
}

func NewApiResponse(description string, resultCode int64, errCode interface{}) *ApiResponse {
	return &ApiResponse{
		ResultCode:  resultCode,
		Description: description,
		ErrorCode:   errCode,
	}
}
