package flags

import "flag"

type Option struct {
	DB   bool   // 初始化数据库
	Load string // 导入数据库文件
	Port int    // 更改端口
}

func Parse() (option *Option) {
	option = new(Option)
	flag.BoolVar(&option.DB, "db", false, "是否初始化数据库")
	flag.StringVar(&option.Load, "load", "", "导入sql数据库文件")
	flag.IntVar(&option.Port, "port", 0, "更改程序运行端口")
	flag.Parse()
	return option
}

// 根据参数运行不同脚本
func (option Option)Run() bool {
	if option.DB {
		DB()
		return true
	}
	if option.Load != "" {
		Load()
		return true
	}
	if option.Port != 0 {
		Port(option.Port)
	}
	return false
}