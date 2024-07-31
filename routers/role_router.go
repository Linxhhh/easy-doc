package routers

import (
	"github.com/Linxhhh/easy-doc/middleware"
	"github.com/Linxhhh/easy-doc/api"
)

func (router RouterGroup) RoleRouter() {
	app := api.App.RoleApi
	router.POST("roles", middleware.JwtAdmin, app.RoleCreate)   // 角色创建（admin）
	router.PUT("roles", middleware.JwtAdmin, app.RoleUpdate)    // 角色更新（admin）
	router.GET("roles", middleware.JwtAdmin, app.RoleList)      // 角色列表（admin）
	router.DELETE("roles", middleware.JwtAdmin, app.RoleRemove) // 角色删除（admin）
}
