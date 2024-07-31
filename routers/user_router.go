package routers

import (
	"github.com/Linxhhh/easy-doc/middleware"
	"github.com/Linxhhh/easy-doc/api"
)

func (router RouterGroup) UserRouter() {
	app := api.App.UserApi
	router.POST("login", app.UserLogin)                      // 登录用户
	router.GET("logout", middleware.JwtAuth, app.UserLogout) // 注销用户（检查是否登录）

	router.POST("users", middleware.JwtAdmin, app.UserCreate)   // 创建用户（admin）
	router.PUT("users", middleware.JwtAdmin, app.UserUpdate)    // 更新用户（admin）
	router.GET("users", middleware.JwtAdmin, app.UserList_v2)   // 用户列表（admin）
	router.DELETE("users", middleware.JwtAdmin, app.UserRemove) // 删除用户（admin）

	router.GET("user_info", middleware.JwtAuth, app.UserInfo)       // 用户信息
	router.PUT("user_info", middleware.JwtAuth, app.UserInfoUpdate) // 修改信息
	router.PUT("user_pwd", middleware.JwtAuth, app.UserPwdUpdate)   // 修改密码
}
