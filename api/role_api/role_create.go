package role_api

import (
	"time"

	"github.com/Linxhhh/easy-doc/models"
	"github.com/Linxhhh/easy-doc/service/common/res"
	"github.com/Linxhhh/easy-doc/global"

	"github.com/gin-gonic/gin"
)

type RoleCreateRequest struct {
	RoleName string `json:"roleName" binding:"required,min=2,max=16"`
}

func (RoleApi) RoleCreate(ctx *gin.Context) {

	var request RoleCreateRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		res.FailWithMsg("参数错误！", ctx)
		return
	}

	var role models.RoleModel
	err = global.DB.Take(&role, "roleName = ?", request.RoleName).Error
	if err == nil {
		res.FailWithMsg("角色名称已存在", ctx)
		return
	}

	now := time.Now()
	global.DB.Create(&models.RoleModel{
		Model: models.Model{
			CreateAt: now,
			UpdateAt: now,
		},
		RoleName: request.RoleName,
	})
	res.OKWithMsg("角色创建成功", ctx)
}
