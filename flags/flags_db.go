package flags

import (
	"github.com/Linxhhh/easy-doc/models"
	"github.com/Linxhhh/easy-doc/global"

	"github.com/sirupsen/logrus"
)

/*
在 GORM 中，迁移是指根据定义的模型表结构，自动创建或更新数据库表结构的过程。
在此处，使用 GORM 的 AutoMigrate 方法来执行迁移操作。
*/
func DB() {
	err := global.DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&models.UserModel{},    // 用户表
		&models.RoleModel{},    // 角色表
		&models.DocModel{},     // 文档表
		&models.RoleDocModel{}, // 角色文档表
		&models.ImageModel{},   // 图片表
		&models.LoginModel{},   // 登录表
		&models.DocDataModel{}, // 文档数据表
		&models.LogModel{},     // 日志表
	)
	if err != nil {
		logrus.Fatalf("数据库迁移失败，%s", err.Error())
	}
	logrus.Info("数据库迁移成功")
}
