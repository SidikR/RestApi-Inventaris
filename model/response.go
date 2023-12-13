package model

type ApiResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
