package jwts

import (
	"github.com/Linxhhh/easy-doc/global"

	"github.com/dgrijalva/jwt-go/v4"
)

// 解析用户 token
func ParseToken(token string) (*CustomClaims, error) {
	
	Token, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(global.Config.Jwt.Serect), nil
	})
	if err != nil {
		// 解析 token 异常
		return nil, err
	}
	if !Token.Valid {
		// 令牌无效
		return nil, &jwt.TokenNotValidYetError{}
	}

	claims, ok := Token.Claims.(*CustomClaims)
	if !ok {
		// 数据不一致
		return nil, &jwt.InvalidClaimsError{}
	}

	return claims, nil
}