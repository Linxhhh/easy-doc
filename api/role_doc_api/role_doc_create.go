package role_doc_api

import (
	"time"

	"github.com/Linxhhh/easy-doc/models"
	"github.com/Linxhhh/easy-doc/service/common/res"
	"github.com/Linxhhh/easy-doc/global"

	"github.com/gin-gonic/gin"
)

type RoleDocCreateRequest struct {
	RoleId uint `json:"roleId" binding:"required" label:"角色id"`
	DocId  uint `json:"docId" binding:"required" label:"文档id"`
}

func (RoleDocApi) RoleDocCreate(ctx *gin.Context) {

	var request RoleDocCreateRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		res.FailWithMsg("参数错误！", ctx)
		return
	}

	var roleDoc models.RoleDocModel
	err = global.DB.Take(&roleDoc, "role_id = ? and doc_id = ?", request.RoleId, request.DocId).Error
	if err == nil {
		res.FailWithMsg("该记录已存在！", ctx)
		return
	}

	now := time.Now()
	global.DB.Create(&models.RoleDocModel{
		Model: models.Model{
			CreateAt: now,
			UpdateAt: now,
		},
		RoleId: request.RoleId,
		DocId:  request.DocId,
	})

	res.OKWithMsg("添加成功", ctx)
}
