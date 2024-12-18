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

// @Summary say hello
// @Description only for example
// @ID ExampleHello
// @Produce json
// @Param authorization header string true "用户ID 会话令牌" example(50123 42e782dc8e131d6989dc772c2c3e87a3)
// @Param game_id query int true "游戏ID" example(1001)
// @Success 200 {object} Response{data=ZonesResponse} "成功"
// @Failure 400 {object} Response "错误请求"
// @Failure 401 {object} Response "未授权"
// @Router /module/line/zones [get]
func (e *Example) Get(ctx *gin.Context) {
	// id, err := strconv.Atoi(ctx.Query("id"))

}
