package models

type RoleModel struct {
	Model
	RoleName string      `gorm:"column:roleName;size:16;comment:角色名称;not null" json:"roleName"`              // 角色名称
	IsSystem bool        `gorm:"column:isSystem;comment:系统角色;not null" json:"isSystem"`                      // 是否为系统角色，不可删
	DocList  []DocModel  `gorm:"many2many:role_doc_models;joinForeignKey:RoleId;joinReferences:DocId" josn:"-"` // 角色拥有的文档列表
	UserList []UserModel `gorm:"ForeignKey:RoleId" josn:"-"`                                                    // 为该角色的用户列表
}

type RoleIdRequest struct {
	ID uint `uri:"id" json:"id" binding:"required"`
}

/*
	记录一个 BUG ：
	go 的规范格式是小驼峰，而 gorm 的规范格式是下划线
	gorm 在预加载模型时，如果是多对多的关系，则会出现命名冲突！

	则应该修改 role_doc_models 中的列名为 role_id 和 doc_id
*/
