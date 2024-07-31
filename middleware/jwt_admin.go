package middleware

import (
	"github.com/Linxhhh/easy-doc/service/redis_service"
	"github.com/Linxhhh/easy-doc/utils/jwts"
	"github.com/Linxhhh/easy-doc/service/common/res"

	"github.com/gin-gonic/gin"
)

/*
	超级管理员验证
	检验用户是否是超级管理员，应用于 “以超级管理员登录的场景”.
*/

func JwtAdmin(ctx *gin.Context) {
	token := ctx.Request.Header.Get("token")
	if token == "" {
		res.FailWithMsg("未携带令牌!", ctx)
		ctx.Abort()
		return
	}

	claims, err := jwts.ParseToken(token)
	if err != nil {
		res.FailWithMsg("令牌错误!", ctx)
		ctx.Abort()
		return
	}

	ok := redis_service.CheckLogout(token)
	if ok {
		res.FailWithMsg("token已经注销!", ctx)
		ctx.Abort()
		return
	}

	// 判断角色是否为超级管理员
	if claims.RoleId != 3 {
		res.FailWithMsg("权限错误!", ctx)
		ctx.Abort()
		return
	}

	// 如果通过验证，则设置 claims 上下文
	ctx.Set("claims", claims)
}
