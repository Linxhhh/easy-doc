package doc_api

import (
	"fmt"

	"github.com/Linxhhh/easy-doc/models"
	"github.com/Linxhhh/easy-doc/service/common/res"
	"github.com/Linxhhh/easy-doc/global"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (DocApi) DocRemove(ctx *gin.Context) {

	var idRequest models.DocIdRequest
	err := ctx.ShouldBindJSON(&idRequest)
	if err != nil {
		res.FailWithMsg("参数错误！", ctx)
		return
	}

	var doc models.DocModel
	err = global.DB.Take(&doc, idRequest.ID).Error
	if err != nil {
		res.FailWithMsg("该文档不存在！", ctx)
		return
	}

	// 获取子文档的 ID 列表
	childList := models.FindAllChildDocList(doc)
	docList := append(childList, doc)

	var docIdList []uint
	for _, model := range docList {
		docIdList = append(docIdList, model.ID)
	}

	// 使用事务，来保证数据一致性
	err = DocRemoveService(docList, docIdList)
	if err != nil {
		global.Log.Errorf("删除文档 %s 失败，err is %s", doc.Title, err.Error())
		res.FailWithMsg("删除文档失败！", ctx)
		return
	}
	res.OKWithMsg(fmt.Sprintf("删除文档成功 共删除 %d 篇文档", len(docList)), ctx)
}

func DocRemoveService(docList []models.DocModel, docIdList []uint) (err error) {

	err = global.DB.Transaction(func(tx *gorm.DB) error {

		// 删除角色-文档表
		var roleDocList []models.RoleDocModel
		err = global.DB.Find(&roleDocList, "doc_id in ?", docIdList).Delete(&roleDocList).Error
		if err != nil {
			return err
		}

		// 删除文档
		err = global.DB.Delete(&docList).Error
		if err != nil {
			return err
		}
		return nil
	})
	return err
}
