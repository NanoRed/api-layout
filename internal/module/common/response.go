package common

type Response struct {
	Code Code   `json:"code" binding:"required" example:"0"`
	Info string `json:"info" binding:"required" example:"success"`
	Data any    `json:"data" extensions:"x-nullable"`
}
