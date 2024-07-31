package testdata

import (
	"fmt"
	"testing"

	"github.com/Linxhhh/easy-doc/global"
	"github.com/Linxhhh/easy-doc/utils/jwts"
	"github.com/Linxhhh/easy-doc/core"
)

/*
在测试时，需要修改 init_config.go 中的 yaml 文件路径！
将 const yamlPath = "settings.yaml" 修改为 const yamlPath = "../settings.yaml"
*/
func TestGenAndPraseToken(t *testing.T) {

	global.Log = core.InitLogger()
	global.Config = core.InitConfig()

	token, err := jwts.GenToken(jwts.JwtPayload{
		UserName: "lxh",
		UserId:   1,
		RoleId:   2,
	})
	fmt.Println(token, err)
	fmt.Println(jwts.ParseToken(token))
}
