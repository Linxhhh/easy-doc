package config

import "fmt"

type Mysql struct {
	Host         string `yaml:"host"`
	Port         int    `yaml:"port"`
	Config       string `yaml:"config"`         // 高级配置
	DB           string `yaml:"db"`             // 数据库名
	Username     string `yaml:"username"`       // 数据库用户
	Password     string `yaml:"password"`       // 数据库密码
	LogLevel     string `yaml:"loglevel"`       // 是否开启 Gorm 全局日志
}

func (m *Mysql) Dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", m.Username, m.Password, m.Host, m.Port, m.DB, m.Config)
}