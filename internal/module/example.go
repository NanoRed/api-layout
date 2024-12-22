package module

import (
	"api-layout/internal/model"
	"api-layout/internal/module/common"
	"api-layout/internal/service"
	"api-layout/pkg/logger"
	"strconv"

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

func (e *Example) Attach(router *gin.Engine) {
	example := router.Group("/api/module/example")
	{
		example.GET("/data", e.Get)
		example.PUT("/data", e.Put)
	}
}

// @Summary Get an example
// @Description Get an example by id
// @ID get-an-example
// @Produce json
// @Param id query int true "example id" example(10001)
// @Success 200 {object} common.Response{payload=model.Example} "ok"
// @Failure 400 {object} common.Response "bad request"
// @Failure 500 {object} common.Response "internal server error"
// @Router /module/example/data [get]
func (e *Example) Get(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Query("id"), 10, 64)
	if err != nil {
		logger.Error("failed to parse id: %v", err)
		common.CodeIllegal.Response(ctx, nil)
		return
	}
	data, err := e.svcExample.Find(ctx, id)
	if err != nil {
		logger.Error("failed to find example: %v", err)
		common.CodeInternal.Response(ctx, nil)
		return
	}
	common.CodeSuccess.Response(ctx, data)
}

// @Summary Put an example
// @Description Put an example
// @ID put-an-example
// @Accept json
// @Produce json
// @Param data body model.Example true "data to be saved"
// @Success 200 {object} common.Response "ok"
// @Failure 400 {object} common.Response "bad request"
// @Failure 500 {object} common.Response "internal server error"
// @Router /module/example/data [put]
func (e *Example) Put(ctx *gin.Context) {
	var data model.Example
	if err := ctx.ShouldBindJSON(&data); err != nil {
		logger.Error("failed to bind json: %v", err)
		common.CodeIllegal.Response(ctx, nil)
		return
	}
	if err := e.svcExample.Save(ctx, &data); err != nil {
		logger.Error("failed to save example: %v", err)
		common.CodeInternal.Response(ctx, nil)
		return
	}
	common.CodeSuccess.Response(ctx, nil)
}
