package user_api

import (
	"time"

	"github.com/Linxhhh/easy-doc/models"
	"github.com/Linxhhh/easy-doc/service/common/res"
	"github.com/Linxhhh/easy-doc/utils/pwd_hash"
	"github.com/Linxhhh/easy-doc/global"

	"github.com/gin-gonic/gin"
)

type UserCreateRequest struct {
	UserName string `json:"userName" binding:"required"`
	Password string `json:"password" binding:"required"`
	// RoleId   uint   `json:"roleId"`
}

func (UserApi) UserCreate(ctx *gin.Context) {

	var request UserCreateRequest

	// 解析请求参数
	err := ctx.BindJSON(&request)
	if err != nil {
		res.FailWithMsg("信息填写不完整！", ctx)
		return
	}

	// 检查用户名是否已经存在
	var user models.UserModel
	err = global.DB.Take(&user, "userName = ?", request.UserName).Error
	if err == nil {
		res.FailWithMsg("用户名重复！", ctx)
		return
	}

	// 创建用户
	now := time.Now()
	newUser := models.UserModel{
		Model: models.Model{
			CreateAt: now,
			UpdateAt: now,
		},
		UserName: request.UserName,
		Password: pwd_hash.HashPwd(request.Password), // 将用户输入的密码进行哈希加密，再存入数据库
		IP:       ctx.ClientIP(),                     // 使用 ClientIP 获取客户端 IP
		RoleId:   1,
	}
	err = global.DB.Create(&newUser).Error
	if err != nil {
		res.FailWithMsg("用户创建失败！", ctx)
		return
	}
	res.OKWithMsg("用户创建成功！", ctx)
}
