package models

type UserModel struct {
	Model
	UserName string `gorm:"column:userName;size:36;comment:用户名;unique;not null" json:"userName"`  // 用户名
	Password string `gorm:"column:password;size:128;comment:用户密码" json:"-"`                      // 用户密码
	Avatar   string `gorm:"column:avatar;size:256;comment:用户头像" json:"avatar"`                   // 用户头像
	Email    string `gorm:"column:email;size:128;comment:电子邮箱" json:"email"`                     // 电子邮箱
	Addr     string `gorm:"column:addr;size:64;comment:居住地址" json:"addr"`                        // 居住地址
	Token    string `gorm:"column:token;size:64;comment:身份令牌" json:"-"`                          // 身份令牌
	IP       string `gorm:"column:ip;size:16;comment:IP地址" json:"ip"`                              // IP

	RoleId    uint      `gorm:"column:roleId;comment:用户的角色" json:"roleId"` // 用户角色
	RoleModel RoleModel `gorm:"foreignKey:RoleId" json:"roleModel"`
}
