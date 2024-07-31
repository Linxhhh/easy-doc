package routers

import (
	"github.com/Linxhhh/easy-doc/middleware"
	"github.com/Linxhhh/easy-doc/api"
)

func (router RouterGroup) RoleDocRouter() {
	app := api.App.RoleDocApi
	router.GET("role_doc/:id", middleware.JwtAdmin, app.RoleDocList)  // 角色-文档列表（admin）
	router.POST("role_doc", middleware.JwtAdmin, app.RoleDocCreate)   // 添加角色-文档（admin）
	router.DELETE("role_doc", middleware.JwtAdmin, app.RoleDocRemove) // 删除角色-文档（admin）

	router.GET("role_doc", app.RoleDocTree) // 角色文档树，不同角色在前端页面看到的文档目录不同
}
