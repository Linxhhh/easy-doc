package user_api

import (
	"fmt"
	"os"
	"path"

	"github.com/Linxhhh/easy-doc/models"
	"github.com/Linxhhh/easy-doc/service/common/res"
	"github.com/Linxhhh/easy-doc/global"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type IDListRequest struct {
	IDList []uint `json:"idList" form:"idList" binding:"required"`
}

func (UserApi) UserRemove(ctx *gin.Context) {

	var request IDListRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		res.FailWithMsg("参数错误！", ctx)
		return
	}

	// 一致性校验
	var userList []models.UserModel
	global.DB.Find(&userList, request.IDList)
	if len(userList) != len(request.IDList) {
		res.FailWithMsg("数据一致性校验不通过！", ctx)
		return
	}

	ok, fail := 0, 0
	for _, user := range userList {
		// global.DB.Delete(&user)
		err := UserRemoveService(user)
		if err != nil {
			logrus.Errorf("删除用户 %d 失败，err is %s", user.ID, err.Error())
			fail++
		} else {
			logrus.Infof("删除用户 %d 成功", user.ID)
			ok++
		}
	}

	if ok > 0 {
		res.OKWithMsg(fmt.Sprintf("%d 位用户成功删除！", ok), ctx)
	}
	if fail > 0 {
		res.FailWithMsg(fmt.Sprintf("%d 位用户删除失败！", fail), ctx)
	}
}

/*
	1. 不添加一致性校验的写法：

	for _, id := range request.IDList {
		global.DB.Delete(&models.UserModel{}, id)
	}

	2. 删除用户时，需要连带删除其数据，比如头像
	   注意，loginModel 不要删除！
*/

func UserRemoveService(user models.UserModel) (err error) {
	// 使用事务，来保证数据一致性
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		// 删除图片目录
		err = os.RemoveAll(path.Join("uploads", user.UserName))
		if err != nil {
			return err
		}
		// 更新图片数据库
		var imgList []models.ImageModel
		tx.Find(&imgList, "userId = ?", user.ID)
		if len(imgList) > 0 {
			err := tx.Delete(&imgList).Error
			if err != nil {
				return err
			}
		}
		// 更新用户数据库
		err = tx.Delete(&user).Error
		if err != nil {
			return err
		}
		return nil
	})
	return err
}
