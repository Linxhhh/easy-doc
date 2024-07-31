package user_api

import (
	"fmt"

	"github.com/Linxhhh/easy-doc/service/common/list"
	"github.com/Linxhhh/easy-doc/service/common/res"
	"github.com/Linxhhh/easy-doc/models"

	"github.com/gin-gonic/gin"
)

func (UserApi) UserList_v2(ctx *gin.Context) {

	// 绑定用户浏览器提供的参数（分页）
	var query list.Querys
	ctx.ShouldBindQuery(&query)

	// 通过 QueryList 来获得用户列表，增加系统自定义参数
	_list, count, _ := list.QueryList(models.UserModel{}, list.Options{
		Querys:  query,
		Likes:   []string{"userName"}, // 可以对多个字段进行模糊匹配，但是用户表只有 UserName，如果有 nickName 的话，则可添加
		Debug:   true,
		PreLoad: []string{"RoleModel"}, // 可以预加载模型，在此次预加载角色模型，用于显示
	})

	res.OK(_list, fmt.Sprintf("获取列表成功！count: %d", count), ctx)
}
