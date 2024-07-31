package redis_service

import (
	"github.com/Linxhhh/easy-doc/global"
	"time"
)

const prefix = "Logout_"

/*
	将 token 缓存，表示该 token 已经注销
*/
func Logout(token string, expiration time.Duration) error {
	err := global.Redis.Set(prefix + token, "", expiration).Err()
	return err
}
/*
	查询缓存，如果存在指定的 token，表明该 token 已经注销
*/
func CheckLogout(token string) bool {
	_, err := global.Redis.Get(prefix + token).Result()
	if err != nil {
		return false
	}
	return true
}