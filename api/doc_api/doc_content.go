package doc_api

import (
	"github.com/Linxhhh/easy-doc/global"
	"github.com/Linxhhh/easy-doc/models"
	"github.com/Linxhhh/easy-doc/service/common/res"
	"github.com/Linxhhh/easy-doc/service/redis_service"
	"github.com/Linxhhh/easy-doc/utils/jwts"

	"github.com/gin-gonic/gin"
)

type DocContentResponse struct {
	Content   string `json:"content"`
	ReadCount int    `json:"readCount"` // 浏览量
	DiggCount int    `json:"diggCount"` // 点赞量
}

func (DocApi) DocContent(ctx *gin.Context) {

	var idRequest models.DocIdRequest
	err := ctx.ShouldBindUri(&idRequest)
	if err != nil {
		res.FailWithMsg("参数错误", ctx)
		return
	}

	// 解析 token，判断用户是否登录
	token := ctx.Request.Header.Get("token")
	claims, err := jwts.ParseToken(token)
	var roleId uint = 4 // 访客
	if err == nil {
		// 如果已经登录
		roleId = claims.RoleId
	}

	// 判断角色对文档的访问权限
	var roleDoc models.RoleDocModel
	err = global.DB.Preload("RoleModel").Take(&roleDoc, "role_id = ? and doc_id = ?", roleId, idRequest.ID).Error
	if err != nil {
		res.FailWithMsg("文档鉴权失败！", ctx)
		return
	}

	doc := roleDoc.DocModel

	docDigg := redis_service.NewDocDigg().GetById(doc.ID)
	redis_service.NewDocRead().SetById(doc.ID)
	docRead := redis_service.NewDocRead().GetById(doc.ID)

	var response = DocContentResponse{
		DiggCount: docDigg + doc.DiggCount,
		ReadCount: docRead + doc.ReadCount,
	}

	res.OKWithData(response, ctx)
}
