package user_api

import (
	"regexp"

	"github.com/Linxhhh/easy-doc/models"
	"github.com/Linxhhh/easy-doc/service/common/res"
	"github.com/Linxhhh/easy-doc/utils/jwts"
	"github.com/Linxhhh/easy-doc/global"

	"github.com/gin-gonic/gin"
)

type InfoUpdateRequest struct {
	Avatar string `json:"avatar"`
	Email  string `json:"email"`
	Addr   string `json:"addr"`
}

func (UserApi) UserInfoUpdate(ctx *gin.Context) {

	var request InfoUpdateRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		res.FailWithMsg("参数错误！", ctx)
		return
	}

	// 获取 claims 上下文
	_claims, _ := ctx.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	var user models.UserModel
	err = global.DB.Take(&user, "ID = ?", claims.UserId).Error
	if err != nil {
		res.FailWithMsg("用户不存在！", ctx)
		return
	}

	// 修改用户头像
	if request.Avatar != "" {
		// 在数据库中，检查头像路径是否存在
		var img models.ImageModel
		err := global.DB.Take(&img, "userId = ? and path = ?", claims.UserId, request.Avatar[1:]).Error
		if err != nil {
			res.FailWithMsg("头像路径不存在！", ctx)
			return
		}
	}

	// 修改用户邮箱
	if request.Email != "" {
		// 检查邮箱格式
		if !IsValidEmail(request.Email) {
			res.FailWithMsg("邮箱格式不正确！", ctx)
			return
		}
		user.Email = request.Email
	}

	// 修改用户住址
	if request.Addr != "" {
		user.Addr = request.Addr
	}

	// 更新用户个人信息
	global.DB.Save(&user)
	res.OKWithMsg("信息修改成功！", ctx)
}

/*
	不推荐以下这种写法，因为不能逐个判断是否为空值

	global.DB.Model(&user).Updates(models.UserModel{
		Avatar: request.Avatar,
		Email: request.Email,
		Addr: request.Addr,
	})
*/

// 正则表达式，检查邮箱格式（QQ邮箱或谷歌邮箱）
func IsValidEmail(email string) bool {
	var emailRegex = regexp.MustCompile(`^(\d+@qq\.com)|(.*@gmail\.com)$`)
	return emailRegex.MatchString(email)
}
