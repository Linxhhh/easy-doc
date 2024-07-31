package routers

import (
	"github.com/Linxhhh/easy-doc/middleware"
	"github.com/Linxhhh/easy-doc/api"
)

func (router RouterGroup) ImageRouter() {
	app := api.App.ImageApi
	router.POST("image", middleware.JwtAuth, app.ImageUpload)    // 图片上传
	router.GET("images", middleware.JwtAuth, app.ImageList)      // 图片列表
	router.DELETE("images", middleware.JwtAuth, app.ImageRemove) // 图片删除
}
