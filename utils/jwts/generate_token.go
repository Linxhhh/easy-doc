package jwts

import (
	"github.com/Linxhhh/easy-doc/global"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
)

// 生成用户 token
func GenToken(user JwtPayload) (string, error) {

	claims := CustomClaims{
		JwtPayload: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(time.Duration(global.Config.Jwt.Expires) * time.Hour)), // 到期时间
			Issuer:    global.Config.Jwt.Issuer,                                                     // 签发人
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)   // 加密算法

	return token.SignedString([]byte(global.Config.Jwt.Serect))  // 使用密钥，生成带签名的JWT
}
