package role_api

import (
	"github.com/Linxhhh/easy-doc/models"
	"github.com/Linxhhh/easy-doc/service/common/res"
	"github.com/Linxhhh/easy-doc/global"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (RoleApi) RoleRemove(ctx *gin.Context) {

	var request models.RoleIdRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		res.FailWithMsg("参数错误！", ctx)
		return
	}

	var role models.RoleModel
	err = global.DB.Preload("DocList").Preload("UserList").Take(&role, request.ID).Error
	if err != nil {
		res.FailWithMsg("该角色不存在！", ctx)
		return
	}

	if role.IsSystem {
		res.FailWithMsg("系统角色，不可删除！", ctx)
		return
	}

	err = RoleRemoveService(role)
	if err != nil {
		global.Log.Errorf("删除角色 %s 失败，err is %s", role.RoleName, err.Error())
		res.FailWithMsg("删除角色失败！", ctx)
		return
	}
	global.Log.Infof("删除角色 %s，关联用户 %d 个", role.RoleName, len(role.UserList))
	res.OKWithMsg("删除角色成功", ctx)
}

func RoleRemoveService(role models.RoleModel) (err error) {
	// 使用事务，来保证数据一致性
	err = global.DB.Transaction(func(tx *gorm.DB) error {

		// 统一修改用户的角色 id 为 1 （普通用户）
		if len(role.UserList) > 0 {
			err = global.DB.Model(&role.UserList).Update("roleId", "1").Error
			if err != nil {
				return err
			}
		}

		// 删除角色关联的文档
		if len(role.DocList) > 0 {
			err = global.DB.Model(&role).Association("DocList").Delete(role.DocList)
			if err != nil {
				return err
			}
		}

		// 删除角色
		err = global.DB.Delete(&role).Error
		if err != nil {
			return err
		}

		return nil
	})
	return err
}
