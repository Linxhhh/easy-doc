package models

import (
	"github.com/Linxhhh/easy-doc/global"
	"sort"
	"strings"
)

type DocModel struct {
	Model
	Title       string      `gorm:"column:title;comment:文档标题" json:"title"`
	Content     string      `gorm:"column:content;comment:文档内容" json:"-"`
	DiggCount   int         `gorm:"column:diggCount;comment:点赞量" json:"diggCount"`
	ReadCount   int         `gorm:"column:readCount;comment:浏览量" json:"readCount"`

	ParentId    *uint       `gorm:"column:parentId;comment:父文档ID" json:"parentId"`
	ParentModel *DocModel   `gorm:"foreignKey:ParentId" json:"-"`
	Child       []*DocModel `gorm:"foreignKey:ParentId" json:"-"`
	
	Key         string      `gorm:"column:key;comment:key;not null" json:"key"`   // 用于排序
}

type DocIdRequest struct {
	ID uint `uri:"id" json:"id" binding:"required"`
}

// 递归获取所有父文档ID
func FindAllParentDocList(doc DocModel, idList *[]uint) {

	*idList = append(*idList, doc.ID)
	if doc.ParentId != nil {
		var parentDoc DocModel
		global.DB.Take(&parentDoc, *doc.ParentId)
		FindAllParentDocList(parentDoc, idList)
	}
}

// 递归获取所有子文档
func FindAllChildDocList(doc DocModel) (docList []DocModel) {
	
	global.DB.Preload("Child").Take(&doc)
	for _, model := range doc.Child {
	  docList = append(docList, *model)
	  docList = append(docList, FindAllChildDocList(*model)...)
	}
	return
}

// 返回文档树
func DocTree(parentID *uint) (docList []*DocModel) {

	// 查找同级文档
	var query = global.DB.Where("")
	if parentID == nil {
	  	query.Where("parentId is null")
	} else {
	  	query.Where("parentId = ?", *parentID)
	}

	// 查找下一级文档
	global.DB.Preload("Child").Where(query).Find(&docList)
	for _, model := range docList {
	 	ChildDocs := DocTree(&model.ID)
	  	model.Child = ChildDocs
	}
	return
}

// 按照 key 中点数进行排序
func SortDoc(docList []*DocModel) (minCount int) {

	if len(docList) == 0 {
	 	return
	}

	sort.Slice(docList, func(i, j int) bool {
	  	count1 := GetPotCount(docList[i])
	  	count2 := GetPotCount(docList[j])
	  	if count1 == count2 {
			// 点数相同，按照文档ID升序
			return docList[i].ID < docList[j].ID
	 	}
	  	return count1 < count2
	})
	
	return GetPotCount(docList[0])
}
  
// 获取 key 的点数
func GetPotCount(doc *DocModel) int {
	return strings.Count(doc.Key, ".")
}

// 文档树的一维化
func TreeByOneDimensional(docList []*DocModel) (list []*DocModel) {
	for _, model := range docList {
	  	list = append(list, model)
	  	list = append(list, TreeByOneDimensional(model.Child)...)
	}
	return
}