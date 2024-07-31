package models

import (
	"time"
)

// 这里参考 gorm.Model
type Model struct {
	ID       uint      `gorm:"primaryKey" json:"id"`             // 主键id
	CreateAt time.Time `gorm:"column:createAt" json:"createAt"`  // 添加时间
	UpdateAt time.Time `gorm:"column:updateAt" json:"updateAt"`  // 修改时间
}