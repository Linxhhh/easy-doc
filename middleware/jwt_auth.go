package middleware

import (
	"log"
	"time"

	"github.com/Linxhhh/easy-doc/service/redis_service"
	"github.com/Linxhhh/easy-doc/utils/jwts"
	"github.com/Linxhhh/easy-doc/service/common/res"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gin-gonic/gin"
)

// 登录验证 —— 检验用户是否登录
func JwtAuth(ctx *gin.Context) {
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

	// 如果剩余有效期小于30分钟，则刷新有效期
	if time.Until(claims.ExpiresAt.Time) < time.Minute*30 {

		claims.ExpiresAt = &jwt.Time{Time: time.Now().Add(time.Hour * 2)} // 刷新有效期：未来两小时

		newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenStr, err := newToken.SignedString([]byte("secret"))
		if err != nil {
			log.Printf("jwt 续约失败，err : %s", err)
		} else {
			ctx.Header("jwt-token", tokenStr)
		}
	}

	// 如果通过验证，则设置 claims 上下文
	ctx.Set("claims", claims)
}
