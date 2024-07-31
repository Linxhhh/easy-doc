package role_doc_api

import (
	"github.com/Linxhhh/easy-doc/models"
	"github.com/Linxhhh/easy-doc/service/common/res"
	"github.com/Linxhhh/easy-doc/global"

	"github.com/gin-gonic/gin"
)

type DocTree struct {
	ID       uint      `json:"key"`
	Title    string    `json:"title"`
	Children []DocTree `json:"children"`
	Show     bool      `json:"show"` // 当前角色是否可以看到该文档
}

type RoleDocListResponse struct {
	List []DocTree `json:"list"`
}

func (RoleDocApi) RoleDocList(ctx *gin.Context) {

	var idRequest models.RoleIdRequest
	err := ctx.ShouldBindUri(&idRequest)
	if err != nil || idRequest.ID == 0 {
		res.FailWithMsg("参数错误！", ctx)
		return
	}

	var roleDocList []models.RoleDocModel
	err = global.DB.Preload("RoleModel").Preload("DocModel").Find(&roleDocList, "role_id = ?", idRequest.ID).Error
	if err != nil {
		res.FailWithMsg("该角色不存在！", ctx)
		return
	}

	// 当前角色拥有的文档，设置为 true
	var docIDMap = map[uint]bool{}
	for _, model := range roleDocList {
		docIDMap[model.DocId] = true
	}

	// 构建当前角色拥有的文档树，并转化成列表
	tree := models.DocTree(nil)
	list := DocTreeTransition(tree, docIDMap)

	res.OKWithData(RoleDocListResponse{
		List: list,
	}, ctx)
}

// 角色文档树转换
func DocTreeTransition(docList []*models.DocModel, docIDMap map[uint]bool) (list []DocTree) {

	for _, model := range docList {
		children := DocTreeTransition(model.Child, docIDMap)
		if children == nil {
			children = make([]DocTree, 0)
		}
		docTree := DocTree{
			ID:       model.ID,
			Title:    model.Title,
			Children: children,
			Show:     docIDMap[model.ID], // map中记录的为true
		}
		list = append(list, docTree)
	}
	return
}
