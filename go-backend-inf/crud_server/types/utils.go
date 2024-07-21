package types

type ApiResponse struct {
	Result      int64  `json:"result"` // 1 : success / -1 : fail
	Description string `json:"description"`
}

func NewApiResponse(description string, result int64) *ApiResponse {
	return &ApiResponse{
		Result:      result,
		Description: description,
	}
}
