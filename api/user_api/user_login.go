package user_api

import (
	"github.com/Linxhhh/easy-doc/models"
	"github.com/Linxhhh/easy-doc/service/common/res"
	"github.com/Linxhhh/easy-doc/utils/jwts"
	"github.com/Linxhhh/easy-doc/utils/pwd_hash"
	"github.com/Linxhhh/easy-doc/global"

	"github.com/gin-gonic/gin"
)

type UserLoginRequest struct {
	UserName string `json:"userName" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (UserApi) UserLogin(ctx *gin.Context) {

	var request UserLoginRequest

	// 解析请求参数
	err := ctx.BindJSON(&request)
	if err != nil {
		res.FailWithMsg("信息填写不完整！", ctx)
		return
	}

	// 检查用户名是否已经存在
	var user models.UserModel
	err = global.DB.Take(&user, "userName = ?", request.UserName).Error
	if err != nil {
		res.FailWithMsg("用户名不存在！", ctx)
		return
	}

	// 检查密码是否正确
	if !pwd_hash.CheckPwd(user.Password, request.Password) {
		res.FailWithMsg("密码错误！", ctx)
		return
	}

	// 生成用户token
	token, err := jwts.GenToken(jwts.JwtPayload{
		UserName: user.UserName,
		UserId:   user.ID,
		RoleId:   user.RoleId,
	})
	if err != nil {
		res.FailWithMsg("生成用户令牌错误！", ctx)
		return
	}

	// 返回用户token
	res.OKWithData(token, ctx)
}
