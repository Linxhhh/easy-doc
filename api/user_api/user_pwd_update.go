package user_api

import (
	"github.com/Linxhhh/easy-doc/models"
	"github.com/Linxhhh/easy-doc/service/common/res"
	"github.com/Linxhhh/easy-doc/utils/jwts"
	"github.com/Linxhhh/easy-doc/utils/pwd_hash"
	"github.com/Linxhhh/easy-doc/global"

	"github.com/gin-gonic/gin"
)

type PwdUpdateRequest struct {
	OldPwd string `json:"oldPwd" binding:"required"`
	NewPwd string `json:"newPwd" binding:"required"`
}

func (UserApi) UserPwdUpdate(ctx *gin.Context) {

	var request PwdUpdateRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		res.FailWithMsg("参数错误！", ctx)
		return
	}

	// 获取 claims 上下文
	_claims, _ := ctx.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	var user models.UserModel
	err = global.DB.Take(&user, "ID = ?", claims.UserId).Error
	if err != nil {
		res.FailWithMsg("用户不存在！", ctx)
		return
	}

	// 检查原密码是否正确
	if !pwd_hash.CheckPwd(user.Password, request.OldPwd) {
		res.FailWithMsg("原密码填写错误！", ctx)
		return
	}

	// 更新用户密码
	user.Password = pwd_hash.HashPwd(request.NewPwd)
	global.DB.Save(&user)
	res.FailWithMsg("密码修改成功！", ctx)
}
