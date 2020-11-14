package libs

import (
	"{{.ModuleName}}/global"
	"github.com/gin-gonic/gin"
)

//BaseController
type BaseController struct {
}

//Output 输出json用
func (c *BaseController) Output(ctx *gin.Context) *Resp {
	return NewResp(ctx, global.SConf.IsProd)
}

//PageParams 分页参数
type PageParams struct {
	Page  int64 `form:"page"`
	Limit int64 `form:"limit"`
}

//GetPageParams 获取分页参数
func (c *BaseController) GetPageParams(ctx *gin.Context) (page, skip, limit int64) {
	var pp = PageParams{
		Page:  0,
		Limit: 0,
	}
	_ = ctx.ShouldBindQuery(&pp)
	page = pp.Page
	if page < 1 {
		page = 1
	}
	limit = pp.Limit
	if limit < 1 {
		limit = 10
	}
	skip = (page - 1) * limit
	return
}
