package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Code uint16

const (
	CodeSuccess      Code = 0x0000
	CodeFailure      Code = 0x0001
	CodeIllegal      Code = 0x0002
	CodeUnauthorized Code = 0x0003
	CodeInternal     Code = 0x0004
)

func (c Code) String() string {
	switch c {
	case CodeSuccess:
		return "request successful"
	case CodeFailure:
		return "request failed"
	case CodeIllegal:
		return "illegal request"
	case CodeUnauthorized:
		return "authentication required"
	case CodeInternal:
		return "internal server error"
	}
	return ""
}

func (c Code) Response(ctx *gin.Context, payload any) {
	var httpStatus int
	switch c {
	case CodeSuccess:
		httpStatus = http.StatusOK
	case CodeUnauthorized:
		httpStatus = http.StatusUnauthorized
	case CodeInternal:
		httpStatus = http.StatusInternalServerError
	default:
		httpStatus = http.StatusBadRequest
	}
	ctx.Abort()
	ctx.JSON(httpStatus, Response{
		Code:    c,
		Message: c.String(),
		Payload: payload,
	})
}
