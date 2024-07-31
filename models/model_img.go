package models

import "fmt"

type ImageModel struct {
	Model
	UserId    uint      `gorm:"column:userId;size:16;comment:用户ID;not null" json:"userId"`
	UserModel UserModel `gorm:"foreignKey:UserId" json:"-"`
	Name      string    `gorm:"column:name;comment:图片名;size:64" json:"name"`
	Path      string    `gorm:"column:path;comment:图片路径;size:128" json:"path"`
	Size      int64     `gorm:"column:size;comment:图片大小，单位字节" json:"size"`
	
	Hash      string    `gorm:"column:hash;comment:图片哈希，避免重复;size:64" json:"hash"`
}

func (image ImageModel) GetPath() string {
	return fmt.Sprintf("/%s", image.Path)
}

/*
	注意：文件的内部存储路径，和外部访问路径不同！
*/
