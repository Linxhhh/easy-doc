package user_api

import (
	"time"

	"github.com/Linxhhh/easy-doc/service/common/res"
	"github.com/Linxhhh/easy-doc/service/redis_service"
	"github.com/Linxhhh/easy-doc/utils/jwts"
	"github.com/Linxhhh/easy-doc/global"

	"github.com/gin-gonic/gin"
)

func (UserApi) UserLogout(ctx *gin.Context) {

	// 获取并解析token
	token := ctx.Request.Header.Get("token")
	claims, _ := jwts.ParseToken(token)

	// 计算剩余过期时间，并缓存 token，表示注销该 token
	expires := claims.ExpiresAt
	diff := expires.Time.Sub(time.Now())
	err := redis_service.Logout(token, diff)
	if err != nil {
		global.Log.Errorf("redis 缓存 token 失败，err is %s", err.Error())
		res.FailWithMsg("用户注销失败！", ctx)
		return
	}

	res.OKWithMsg("用户注销成功！", ctx)
}
