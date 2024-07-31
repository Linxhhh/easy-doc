package models

import "encoding/json"

type Level int

const (
	Info    Level = 1
	Warning Level = 2
	Error   Level = 3
)

type LogType int

const (
	LoginType   LogType = 1
	ActionType  LogType = 2
	RuntimeType LogType = 3
)

type LogModel struct {
	Model
	UserId   uint   `gorm:"column:userId"   json:"userId"`
	UserName string `gorm:"column:userName" json:"userName"`
	IP       string `gorm:"column:ip"       json:"ip"`
	Addr     string `gorm:"column:addr"     json:"addr"`

	Level       Level   `gorm:"column:level" json:"level"`             // 日志等级  1 信息 2 警告 3 报错
	Type        LogType `gorm:"column:type" json:"type"`               // 日志类型  1 登录 2 操作 3 运行
	Title       string  `gorm:"column:title" json:"title"`             // 日志标题
	Content     string  `gorm:"column:content" json:"content"`         // 日志详情
	ServiceName string  `gorm:"column:serviceName" json:"serviceName"` // 服务名称
	Status      bool    `gorm:"column:status" json:"status"`           // 登录状态
}

// Level 转 string
func (level Level) ToString() string {
	switch level {
	case Info:
		return "info"
	case Warning:
		return "warning"
	case Error:
		return "error"
	}
	return ""
}

// Level 转 json
func (level Level) ToJSON() ([]byte, error) {
	return json.Marshal(level.ToString())
}

// LogType 转 string
func (t LogType) ToString() string {
	switch t {
	case LoginType:
		return "loginType"
	case ActionType:
		return "actionType"
	case RuntimeType:
		return "runtimeType"
	}
	return ""
}

// LogType 转 json
func (t LogType) ToJSON() ([]byte, error) {
	return json.Marshal(t.ToString())
}