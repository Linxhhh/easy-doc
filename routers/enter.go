package routers

import "github.com/gin-gonic/gin"

type RouterGroup struct {
	*gin.RouterGroup
}

func Routers() *gin.Engine {
	router := gin.Default()

	// 创建路由分组
	apiGroup := router.Group("/api")

	// api 分组绑定回调函数
	routerGroup := RouterGroup{
		apiGroup,
	}
	routerGroup.UserRouter()
	routerGroup.ImageRouter()
	routerGroup.DocRouter()
	routerGroup.RoleRouter()
	routerGroup.RoleDocRouter()

	// 静态目录服务，可查看上传的图片
	router.Static("/uploads", "uploads")

	return router
}