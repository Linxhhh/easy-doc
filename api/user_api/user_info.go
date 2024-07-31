package user_api

import (
	"github.com/Linxhhh/easy-doc/models"
	"github.com/Linxhhh/easy-doc/service/common/res"
	"github.com/Linxhhh/easy-doc/utils/jwts"
	"github.com/Linxhhh/easy-doc/global"

	"github.com/gin-gonic/gin"
)

type UserInfoResponse struct {
	UserName string `json:"userName"` // 用户名
	Avatar   string `json:"avatar"`   // 用户头像
	Email    string `json:"email"`    // 电子邮箱
	Addr     string `json:"addr"`     // 居住地址
	RoleName string `json:"roleName"` // 角色名称
}

func (UserApi) UserInfo(ctx *gin.Context) {

	// 获取 claims 上下文
	_claims, _ := ctx.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	var user models.UserModel
	err := global.DB.Preload("RoleModel").Take(&user, "ID = ?", claims.UserId).Error
	if err != nil {
		res.FailWithMsg("用户不存在！", ctx)
		return
	}

	info := UserInfoResponse{
		UserName: user.UserName,
		Avatar:   user.Avatar,
		Email:    user.Email,
		Addr:     user.Addr,
		RoleName: user.RoleModel.RoleName,
	}

	res.OKWithData(info, ctx)
}
