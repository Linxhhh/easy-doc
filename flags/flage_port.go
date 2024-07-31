package flags

import "github.com/Linxhhh/easy-doc/global"

func Port(port int) {
	global.Config.System.Port = port
}