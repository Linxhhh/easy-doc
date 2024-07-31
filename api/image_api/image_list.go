package image_api

import (
	"fmt"

	"github.com/Linxhhh/easy-doc/service/common/list"
	"github.com/Linxhhh/easy-doc/service/common/res"
	"github.com/Linxhhh/easy-doc/models"

	"github.com/gin-gonic/gin"
)

type ImageListResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Size int64  `json:"size"`
	Path string `json:"path"`
}

func (ImageApi) ImageList(ctx *gin.Context) {

	var query list.Querys
	ctx.ShouldBindQuery(&query)

	_list, count, _ := list.QueryList(models.ImageModel{}, list.Options{
		Querys: query,
		Likes:  []string{"name"},
	})

	imageList := make([]ImageListResponse, count)
	for i, model := range _list {
		imageList[i] = ImageListResponse{
			ID:   model.ID,
			Name: model.Name,
			Size: model.Size,
			Path: model.GetPath(),
		}
	}

	res.OK(imageList, fmt.Sprintf("获取列表成功！count: %d", count), ctx)
}
