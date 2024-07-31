package role_api

import (
	"github.com/Linxhhh/easy-doc/models"
	"github.com/Linxhhh/easy-doc/service/common/res"
	"github.com/Linxhhh/easy-doc/global"

	"github.com/gin-gonic/gin"
)

type RoleUpdateRequest struct {
	ID       uint   `json:"id" binding:"required"`
	RoleName string `json:"roleName" binding:"required,min=2,max=16"`
}

func (RoleApi) RoleUpdate(ctx *gin.Context) {

	var request RoleUpdateRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		res.FailWithMsg("参数错误！", ctx)
		return
	}

	var role models.RoleModel
	err = global.DB.Take(&role, "id = ?", request.ID).Error
	if err != nil {
		res.FailWithMsg("该角色不存在！", ctx)
		return
	}
	err = global.DB.Take(&role, "roleName = ?", request.RoleName).Error
	if err == nil {
		res.FailWithMsg("角色名称已存在", ctx)
		return
	}

	role.RoleName = request.RoleName

	err = global.DB.Save(&role).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("更新失败", ctx)
		return
	}
	res.OKWithMsg("更新成功", ctx)
}
