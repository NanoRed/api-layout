package module

import (
	"api-layout/internal/service"

	"github.com/gin-gonic/gin"
)

type Example struct {
	svcExample *service.Example
}

func NewExample(svcExample *service.Example) *Example {
	example := &Example{
		svcExample: svcExample,
	}
	return example
}

// @Summary Get an example
// @Description Get an example by id
// @ID get-an-example
// @Accept json
// @Produce json
// @Param authorization header string true "auth token" example(42e782dc8e131d6989dc772c2c3e87a3)
// @Param id query int true "example id" example(10001)
// @Success 200 {object} common.Response{data=model.Example} "success"
// @Failure 400 {object} common.Response "bad request"
// @Failure 401 {object} common.Response "unauthorized"
// @Router /module/line/zones [get]
func (e *Example) Get(ctx *gin.Context) {
	// id, err := strconv.Atoi(ctx.Query("id"))
}
