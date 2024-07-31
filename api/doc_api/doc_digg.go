package doc_api

import (
	"github.com/Linxhhh/easy-doc/service/common/res"
	"github.com/Linxhhh/easy-doc/service/redis_service"
	"github.com/Linxhhh/easy-doc/models"

	"github.com/gin-gonic/gin"
)

func (DocApi) DocDigg(ctx *gin.Context) {

	var idRequest models.DocIdRequest
	err := ctx.ShouldBindUri(&idRequest)
	if err != nil {
		res.FailWithMsg("参数错误！", ctx)
		return
	}

	redis_service.NewDocDigg().SetById(idRequest.ID)
	res.OKWithMsg("文档点赞成功", ctx)
}
