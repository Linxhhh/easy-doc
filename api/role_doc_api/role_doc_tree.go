package role_doc_api

import (
	"github.com/Linxhhh/easy-doc/models"
	"github.com/Linxhhh/easy-doc/service/common/res"
	"github.com/Linxhhh/easy-doc/service/redis_service"
	"github.com/Linxhhh/easy-doc/utils/jwts"
	"github.com/Linxhhh/easy-doc/global"
	utils "github.com/Linxhhh/easy-doc/utils/max_prefix"

	"github.com/gin-gonic/gin"
)

type RoleDocTree struct {
	ID       uint          `json:"key"`
	Title    string        `json:"title"`
	Children []RoleDocTree `json:"children"`
}

type RoleDocTreeResponse struct {
	List []RoleDocTree `json:"list"`
}

func (RoleDocApi) RoleDocTree(ctx *gin.Context) {

	token := ctx.Request.Header.Get("token")
	claims, err := jwts.ParseToken(token)
	ok := redis_service.CheckLogout(token)

	// 用户角色，默认为访客
	var roleID = uint(4)
	if err == nil && !ok {
		roleID = claims.RoleId
	}

	var response = RoleDocTreeResponse{
		List: make([]RoleDocTree, 0),
	}

	// 获取角色拥有的文档ID列表
	var docIDList []uint
	var roleDocList []models.RoleDocModel
	global.DB.Preload("RoleModel").Preload("DocModel").Find(&roleDocList, "role_id = ?", roleID).Select("doc_id").Scan(&docIDList)
	if len(roleDocList) == 0 {
		res.OKWithData(response, ctx)
		return
	}

	// 获取文档列表
	var docList []*models.DocModel
	global.DB.Find(&docList, docIDList)

	// 文档列表排序
	var docListSorted = new([]*models.DocModel)
	minCount := models.SortDoc(docList)
	for _, model := range docList {
		// 根文档
		if models.GetPotCount(model) == minCount {
			*docListSorted = append(*docListSorted, model)
			continue
		}
		// 子文档
		insertDoc(docListSorted, model)
	}

	list := RoleDocTreeTransition(*docListSorted)
	response.List = list
	res.OKWithData(response, ctx)
}

// 角色文档树转换
func RoleDocTreeTransition(docList []*models.DocModel) (list []RoleDocTree) {

	for _, model := range docList {
		children := RoleDocTreeTransition(model.Child)
		if children == nil {
			children = make([]RoleDocTree, 0)
		}
		docTree := RoleDocTree{
			ID:       model.ID,
			Title:    model.Title,
			Children: children,
		}
		list = append(list, docTree)
	}
	return
}

// 找到父文档
func insertDoc(docList *[]*models.DocModel, doc *models.DocModel) {

	// 将根文档树一维化
	oneDimensionalDocList := models.TreeByOneDimensional(*docList)
	var keys []string
	for _, model := range oneDimensionalDocList {
		keys = append(keys, model.Key)
	}

	// 通过最大字符前缀匹配，找到后面最有可能匹配的key
	_, index := utils.FindMaxPrefix(doc.Key, keys)
	if index == -1 {
		*docList = append(*docList, doc)
	} else {
		oneDimensionalDocList[index].Child = append(oneDimensionalDocList[index].Child, doc)
	}

}
