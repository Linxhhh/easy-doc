package models

type LoginModel struct {
	Model
	UserId    uint      `gorm:"column:userId;size:16;comment:用户ID;not null" json:"userId"`
	UserModel UserModel `gorm:"foreignKey:UserId" json:"-"`
	IP        string    `gorm:"column:ip;size:16;comment:IP地址" json:"ip"`        // IP
	UA        string    `gorm:"column:ua;size:16;comment:User Agent" json:"ua"`    // User Agent
	Token     string    `gorm:"column:token;size:64;comment:身份令牌" json:"-"`    // 身份令牌
	Addr      string    `gorm:"column:addr;size:64;comment:居住地址" json:"addr"`  // 居住地址
}
