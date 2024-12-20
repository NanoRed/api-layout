package common

type Response struct {
	Code    Code   `json:"code" binding:"required" example:"0"`
	Message string `json:"message" binding:"required" example:"success"`
	Payload any    `json:"payload" extensions:"x-nullable"`
}
