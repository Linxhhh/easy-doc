package role_api

import (
	"fmt"

	"github.com/Linxhhh/easy-doc/service/common/list"
	"github.com/Linxhhh/easy-doc/service/common/res"
	"github.com/Linxhhh/easy-doc/models"

	"github.com/gin-gonic/gin"
)

type RoleListResponse struct {
	models.RoleModel
	DocCount  int `json:"docCount"`  // 角色拥有的文档数
	UserCount int `json:"userCount"` // 角色下的用户数
}

func (RoleApi) RoleList(ctx *gin.Context) {

	var query list.Querys
	ctx.ShouldBindQuery(&query)

	_list, count, _ := list.QueryList(models.RoleModel{}, list.Options{
		Querys:  query,
		Likes:   []string{"roleName"},
		PreLoad: []string{"DocList", "UserList"},
	})

	roleList := make([]RoleListResponse, count)
	for i, model := range _list {
		roleList[i] = RoleListResponse{
			RoleModel: model,
			DocCount:  len(model.DocList),
			UserCount: len(model.UserList),
		}
	}
	res.OK(roleList, fmt.Sprintf("获取列表成功！count: %d", count), ctx)
}
