package doc_api

import (
	"github.com/Linxhhh/easy-doc/models"
	"github.com/Linxhhh/easy-doc/service/common/res"
	"github.com/Linxhhh/easy-doc/service/redis_service"
	"github.com/Linxhhh/easy-doc/global"

	"github.com/gin-gonic/gin"
)

type DocInfoResponse struct {
	ID            uint   `json:"id"`
	Title         string `json:"title"`
	ContentLength int    `json:"contentLength"`
	DiggCount     int    `json:"diggCount"`
	ReadCount     int    `json:"readCount"`
	Key           string `json:"key"`
}

func (DocApi) DocInfo(ctx *gin.Context) {

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

	docDigg := redis_service.NewDocDigg().GetById(doc.ID)
	docRead := redis_service.NewDocRead().GetById(doc.ID)

	var response = DocInfoResponse{
		ID:            doc.ID,
		Title:         doc.Title,
		ContentLength: len(doc.Content),
		DiggCount:     doc.DiggCount + docDigg,
		ReadCount:     doc.ReadCount + docRead,
		Key:           doc.Key,
	}

	res.OK(response, "获取文档信息成功！", ctx)
}
