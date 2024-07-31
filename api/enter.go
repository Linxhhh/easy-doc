package api

import (
	"github.com/Linxhhh/easy-doc/api/image_api"
	"github.com/Linxhhh/easy-doc/api/role_api"
	"github.com/Linxhhh/easy-doc/api/role_doc_api"
	"github.com/Linxhhh/easy-doc/api/user_api"
	"github.com/Linxhhh/easy-doc/api/doc_api"
)

type Api struct {
	UserApi    user_api.UserApi
	ImageApi   image_api.ImageApi
	RoleApi    role_api.RoleApi
	DocApi     doc_api.DocApi
	RoleDocApi role_doc_api.RoleDocApi
}

var App = new(Api)
