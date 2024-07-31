package doc_api

import (
	"fmt"
	"time"

	"github.com/Linxhhh/easy-doc/global"
	"github.com/Linxhhh/easy-doc/models"
	"github.com/Linxhhh/easy-doc/service/common/res"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type DocCreateRequest struct {
	Title    string `json:"title" binding:"required" label:"文章标题"`
	Content  string `json:"content" binding:"required" label:"文章内容"`
	ParentId *uint  `json:"parentId"`
}

func (DocApi) DocCreate(ctx *gin.Context) {

	var request DocCreateRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		res.FailWithMsg("参数错误！", ctx)
		return
	}

	// 判断 ParentId 是否合法
	if request.ParentId != nil {
		if *request.ParentId <= 0 {
			res.FailWithMsg("父文档ID非法！", ctx)
			return
		}
		var parentDoc models.DocModel
		err = global.DB.Take(&parentDoc, *request.ParentId).Error
		if err != nil {
			res.FailWithMsg("父文档不存在！", ctx)
			return
		}
	}

	now := time.Now()
	var docModel = models.DocModel{
		Model: models.Model{
			CreateAt: now,
			UpdateAt: now,
		},
		Title:    request.Title,
		Content:  request.Content,
		ParentId: request.ParentId,
	}
	err = global.DB.Create(&docModel).Error
	if err != nil {
		logrus.Error(err.Error())
		res.FailWithMsg("文档保存失败", ctx)
		return
	}

	// 获取所有父文档ID，取反，拼接成key值
	var (
		idList []uint
		key    string
	)
	models.FindAllParentDocList(docModel, &idList)
	for i, id := range idList {
		if i == 0 {
			key = fmt.Sprintf("%d", id)
		} else {
			key = fmt.Sprintf("%d", id) + "." + key
		}
	}

	global.DB.Model(&docModel).Update("key", key)

	// 每次创建文档，都更新角色-文档表
	err = global.DB.Create(&models.RoleDocModel{
		Model: models.Model{
			CreateAt: now,
			UpdateAt: now,
		},
		RoleId: 3,
		DocId:  docModel.ID,
	}).Error
	if err != nil {
		res.FailWithMsg("角色-文档更新失败！", ctx)
		return
	}

	res.OK(docModel.ID, "文档添加成功", ctx)
}
