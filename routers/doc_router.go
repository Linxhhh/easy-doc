package routers

import (
	"github.com/Linxhhh/easy-doc/middleware"
	"github.com/Linxhhh/easy-doc/api"
)

func (router RouterGroup) DocRouter() {
	app := api.App.DocApi
	router.POST("doc", middleware.JwtAdmin, app.DocCreate)              // 创建文档（admin）
	router.DELETE("doc", middleware.JwtAdmin, app.DocRemove)            // 删除文档（admin）
	router.PUT("doc/:id", middleware.JwtAdmin, app.DocUpdate)           // 更新文档（admin）
	router.GET("doc/info/:id", middleware.JwtAdmin, app.DocInfo)        // 文档信息（admin）
	router.GET("doc/edit/:id", middleware.JwtAdmin, app.DocEditContent) // 文档内容（admin）

	router.GET("doc/:id", app.DocContent)   // 文档内容
	router.GET("doc/digg/:id", app.DocDigg) // 文档点赞
}
