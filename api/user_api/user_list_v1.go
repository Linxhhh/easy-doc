package user_api

import (
	"fmt"

	"github.com/Linxhhh/easy-doc/models"
	"github.com/Linxhhh/easy-doc/service/common/res"
	"github.com/Linxhhh/easy-doc/global"

	"github.com/gin-gonic/gin"
)

type UserListRequest struct {
	Page  int    `json:"page"  form:"page"`
	Limit int    `json:"limit" form:"limit"`
	Key   string `josn:"key"   form:"key"` // 进行模糊匹配
}

// 查询用户列表
func (UserApi) UserList_v1(ctx *gin.Context) {

	// 绑定查询参数
	var request UserListRequest
	ctx.ShouldBindQuery(&request)
	if request.Limit <= 0 {
		request.Limit = 10
	}

	// 分页查询
	offset := (request.Page - 1) * request.Limit

	// 模糊查询
	query := global.DB.Where("userName like ?", fmt.Sprintf("%%%s%%", request.Key))

	var users []models.UserModel
	err := global.DB.Where(query).Limit(request.Limit).Offset(offset).Find(&users).Error
	if err != nil {
		res.FailWithMsg("获取用户列表失败", ctx)
		return
	}
	res.OK(users, "获取用户列表成功", ctx)
}

/*
	1. 如果不携带任何参数，则是获取全部的用户列表，相当于以下代码：

	var users []models.UserModel
	err := global.DB.Find(&users).Error
	if err != nil {
		res.FailWithMsg("获取用户列表失败", ctx)
		return
	}
	res.OK(users, "获取用户列表成功", ctx)

	2. 使用模糊查询时，不要直接采用字符串拼接，防止sql注入，下面为错误示例：

	query := global.DB.Where(fmt.Sprintf("userName like %%%s%%", request.Key))
*/
