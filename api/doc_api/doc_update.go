package doc_api

import (
	"github.com/Linxhhh/easy-doc/models"
	"github.com/Linxhhh/easy-doc/service/common/res"
	"github.com/Linxhhh/easy-doc/global"

	"github.com/gin-gonic/gin"
)

type DocUpdateRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (DocApi) DocUpdate(ctx *gin.Context) {

	var idRequest models.DocIdRequest
	err := ctx.ShouldBindUri(&idRequest)
	if err != nil {
		res.FailWithMsg("参数错误", ctx)
		return
	}

	var request DocUpdateRequest
	err = ctx.ShouldBindJSON(&request)
	if err != nil {
		res.FailWithMsg("参数错误！", ctx)
		return
	}

	var doc models.DocModel
	err = global.DB.Take(&doc, idRequest.ID).Error
	if err != nil {
		res.FailWithMsg("文档不存在", ctx)
		return
	}

	if request.Title == "" && request.Content == "" {
		res.OKWithMsg("文档更新成功！", ctx)
		return
	}

	err = global.DB.Model(&doc).Updates(models.DocModel{
		Title:   request.Title,
		Content: request.Content,
	}).Error
	if err != nil {
		res.FailWithMsg("文档更新失败！", ctx)
		return
	}
	res.OKWithMsg("文档更新成功！", ctx)
}
