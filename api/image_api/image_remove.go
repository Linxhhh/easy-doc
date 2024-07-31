package image_api

import (
	"fmt"
	"os"

	"github.com/Linxhhh/easy-doc/models"
	"github.com/Linxhhh/easy-doc/service/common/res"
	"github.com/Linxhhh/easy-doc/utils/jwts"
	"github.com/Linxhhh/easy-doc/global"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type IDListRequest struct {
	IDList []uint `json:"idList" form:"idList" binding:"required"`
}

func (ImageApi) ImageRemove(ctx *gin.Context) {

	var request IDListRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		res.FailWithMsg("参数错误！", ctx)
		return
	}

	// 获取 claims 上下文
	_claims, _ := ctx.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	// 一致性校验
	var imageList []models.ImageModel
	global.DB.Find(&imageList, "userId = ? and id = ?", claims.UserId, request.IDList)
	if len(imageList) != len(request.IDList) {
		res.FailWithMsg("数据一致性校验不通过！", ctx)
		return
	}

	for _, img := range imageList {
		global.DB.Take(&img, "userId = ? and ID = ?", img.UserId)
	}

	ok, fail := 0, 0
	for _, img := range imageList {
		// global.DB.Delete(&img)
		err := ImageRemoveService(img)
		if err != nil {
			global.Log.Errorf("删除图片 %s 失败，err is %s", img.Path, err.Error())
			fail++
		} else {
			global.Log.Infof("删除用户 %s 成功", img.Path)
			ok++
		}
	}

	if ok > 0 {
		res.OKWithMsg(fmt.Sprintf("%d 张图片成功删除！", ok), ctx)
	}
	if fail > 0 {
		res.FailWithMsg(fmt.Sprintf("%d 张图片删除失败！", fail), ctx)
	}
}

func ImageRemoveService(img models.ImageModel) (err error) {
	// 使用事务，来保证数据一致性
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		// 删除图片
		err = os.Remove(img.Path)
		if err != nil {
			return err
		}
		// 更新数据库
		err = global.DB.Delete(img).Error
		if err != nil {
			return err
		}
		return nil
	})
	return err
}
