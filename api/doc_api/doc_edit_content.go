package doc_api

import (
	"github.com/Linxhhh/easy-doc/models"
	"github.com/Linxhhh/easy-doc/service/common/res"
	"github.com/Linxhhh/easy-doc/global"

	"github.com/gin-gonic/gin"
)

/*
	获取完整的正文，用于编辑文档
*/

func (DocApi) DocEditContent(ctx *gin.Context) {

	var idRequest models.DocIdRequest
	err := ctx.ShouldBindUri(&idRequest)
	if err != nil {
		res.FailWithMsg("参数错误", ctx)
		return
	}

	var doc models.DocModel
	err = global.DB.Take(&doc, idRequest.ID).Error
	if err != nil {
		res.FailWithMsg("该文档不存在！", ctx)
		return
	}
	res.OKWithData(doc.Content, ctx)
}
