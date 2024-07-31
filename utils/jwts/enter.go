package jwts

import "github.com/dgrijalva/jwt-go/v4"

type JwtPayload struct {
	UserName string `json:"userName"`
	UserId   uint   `json:"userId"`
	RoleId   uint   `json:"roleId"`
}

type CustomClaims struct {
	JwtPayload
	jwt.StandardClaims // 标准声明结构体
}
