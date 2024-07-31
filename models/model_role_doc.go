package models

type RoleDocModel struct {
	Model
	RoleId      uint      `gorm:"column:role_id;comment:角色ID" json:"roleId"`
	RoleModel   RoleModel `gorm:"foreignKey:RoleId" json:"-"`
	DocId       uint      `gorm:"column:doc_id;comment:文档ID" json:"docId"`
	DocModel    DocModel  `gorm:"foreignKey:DocId" json:"-"`

	Sort        int       `gorm:"column:sort;comment:排序" json:"sort"`
}