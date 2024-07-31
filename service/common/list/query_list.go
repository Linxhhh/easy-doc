package list

import (
	"github.com/Linxhhh/easy-doc/global"
	"fmt"

	"gorm.io/gorm"
)

type Querys struct {
	Page  int    `json:"page" form:"page"`
	Limit int    `json:"limit" form:"limit"`
	Key   string `json:"key" form:"key"`
	Sort  string `json:"sort" form:"sort"`
}

type Options struct {
	Querys
	Likes   []string // 模糊匹配的字段列表
	Debug   bool     // 是否打印日志
	Where   *gorm.DB // 更加精确的条件查询
	PreLoad []string // 预加载
}

/*
	model：查询某一张表
*/

func QueryList[T any](model T, options Options) (list []T, count int, err error) {
	db := global.DB.Where(model)
	
	if options.Debug {
		db = db.Debug()                  // 开启 Debug 模式
	}

	if options.Sort == "" {
		options.Sort = "createAt desc"   // 设置默认排序
	}

	if options.Limit == 0 {
		options.Limit = 10               // 设置默认 Limit
	}

	if options.Where != nil {
		db.Where(options.Where)          // 设置额外的查询条件
	}

	if options.Key != "" {               // 开启模糊匹配
		likeQuery := global.DB.Where("")
		for index, like := range options.Likes {
			if index == 0 {
				likeQuery = likeQuery.Where(fmt.Sprintf("%s like ?", like), fmt.Sprintf("%%%s%%", options.Key))
			} else {
				likeQuery = likeQuery.Or(fmt.Sprintf("%s like ?", like), fmt.Sprintf("%%%s%%", options.Key))
			}
		}
		db = db.Where(likeQuery)
	}

	count = int(db.Find(&list).RowsAffected)

	for _, s := range options.PreLoad {           // 预加载
		db = db.Preload(s)
	}

	offset := (options.Page - 1) * options.Limit  // 分页

	err = db.Limit(options.Limit).Offset(offset).Order(options.Sort).Find(&list).Error

	return 
}