package user_api

import (
	"github.com/Linxhhh/easy-doc/models"
	"github.com/Linxhhh/easy-doc/service/common/res"
	"github.com/Linxhhh/easy-doc/utils/pwd_hash"
	"github.com/Linxhhh/easy-doc/global"

	"github.com/gin-gonic/gin"
)

/*
	超级管理员权限：可以修改用户的角色、用户名、用户密码。
*/

type UserUpdateRequest struct {
	UserId   uint   `josn:"userId" binding:"required"`
	RoleId   uint   `json:"roleId"`
	UserName string `json:"userName"`
	Password string `json:"password"`
}

func (UserApi) UserUpdate(ctx *gin.Context) {

	var request UserUpdateRequest

	// 解析请求参数
	err := ctx.BindJSON(&request)
	if err != nil {
		res.FailWithMsg("参数错误！", ctx)
		return
	}

	// 检查用户是否已经存在
	var user models.UserModel
	err = global.DB.Take(&user, "ID = ?", request.UserId).Error
	if err != nil {
		res.FailWithMsg("用户不存在！", ctx)
		return
	}

	// 更新用户的角色
	if request.RoleId != 0 {
		// 检查角色是否已经存在
		var role models.RoleModel
		err = global.DB.Take(&role, "ID = ?", request.RoleId).Error
		if err != nil {
			res.FailWithMsg("角色不存在！", ctx)
			return
		}
		user.RoleId = request.RoleId
	}

	// 更新用户的用户名
	if request.UserName != "" {
		// 检查用户名是否重复
		var sameNameUser models.UserModel
		err = global.DB.Take(&sameNameUser, "userName = ?", request.UserName).Error
		if err == nil {
			res.FailWithMsg("用户名重复！", ctx)
			return
		}
		user.UserName = request.UserName
	}

	// 更新用户的密码
	if request.Password != "" {
		pwd := pwd_hash.HashPwd(request.Password)
		user.Password = pwd
	}

	// 保存更新后的用户信息到数据库
	err = global.DB.Save(&user).Error
	if err != nil {
		res.FailWithMsg("更新用户信息失败！", ctx)
		return
	}

	// 返回成功响应
	res.OKWithMsg("用户更新成功！", ctx)
}
