package image_api

import (
	"path"
	"strings"
	"time"

	"github.com/Linxhhh/easy-doc/models"
	"github.com/Linxhhh/easy-doc/service/common/res"
	"github.com/Linxhhh/easy-doc/utils/img_hash"
	"github.com/Linxhhh/easy-doc/utils/jwts"
	"github.com/Linxhhh/easy-doc/global"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

/*
	注意：
	1. 图片、文件的上传是通过 form 表单参数！
	2. 需要对上传的文件进行仔细的安全校验（白名单判断、文件大小判断、重复名称判断、文件哈希判断）
*/

func (ImageApi) ImageUpload(ctx *gin.Context) {

	fileHeader, err := ctx.FormFile("image")
	if err != nil {
		res.FailWithMsg("参数错误！", ctx)
		return
	}

	// 获取 claims 上下文
	_claims, _ := ctx.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	// 白名单判断
	if !ImageWhiteListCheck(fileHeader.Filename) {
		res.FailWithMsg("不支持该文件类型！", ctx)
		return
	}

	// 文件大小判断  2MB
	if fileHeader.Size > int64(2*1024*1024) {
		res.FailWithMsg("图片大小超限！请上传小于 2MB 的图片", ctx)
		return
	}

	// 重复名称判断
	var imageModel models.ImageModel
	err = global.DB.Take(&imageModel, "userId = ? and name = ?", claims.UserId, fileHeader.Filename).Error
	if err == nil {
		res.FailWithMsg("图片名称重复！", ctx)
		return
	}

	// 文件哈希检查
	file, _ := fileHeader.Open()
	fileHash := img_hash.FileMd5(file)
	err = global.DB.Take(&imageModel, "userId = ? and hash = ?", claims.UserId, fileHash).Error
	if err == nil {
		res.FailWithMsg("已经存在相同图片！", ctx)
		return
	}

	// 创建图片模型
	now := time.Now()
	savePath := path.Join("uploads", claims.UserName, fileHeader.Filename)
	imageModel = models.ImageModel{
		Model: models.Model{
			CreateAt: now,
			UpdateAt: now,
		},
		UserId: claims.UserId,
		Name:   fileHeader.Filename,
		Size:   fileHeader.Size,
		Path:   savePath,
		Hash:   fileHash,
	}

	// 使用事务，保存图片，并写入数据库
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		if err := ctx.SaveUploadedFile(fileHeader, savePath); err != nil {
			return err
		}
		if err := tx.Create(&imageModel).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		global.Log.Errorf("图片 %s 上传失败，err is %s", fileHeader.Filename, err.Error())
		res.FailWithMsg("图片上传失败", ctx)
		return
	}

	res.OK("/"+savePath, "图片上传成功！", ctx)
}

// 白名单
var ImageWhiteList = []string{
	"png",
	"jpg",
	"jpeg",
	"webp",
}

// 白名单判断
func ImageWhiteListCheck(fileName string) bool {

	// 截取文件后缀，并转化为小写
	_list := strings.Split(fileName, ".")
	if len(_list) < 2 {
		return false
	}
	postfix := strings.ToLower(_list[len(_list)-1])

	for _, s := range ImageWhiteList {
		if postfix == s {
			return true
		}
	}
	return false
}

/*
	不使用事务，而是使用简单的回滚操作的做法：

	err = global.DB.Create(&imageModel).Error
	if err != nil {
		global.Log.Errorf("图片 %s 写入数据库失败，err is %s", fileHeader.Filename, err.Error())
		res.FailWithMsg("写入数据库错误", ctx)

		// 写入数据库失败，回滚，删除图片
		err = os.Remove(savePath)
		if err != nil {
			global.Log.Errorf("尝试删除文件 %s 失败，err is %s", savePath, err.Error())
			res.FailWithMsg("写入数据库错误，尝试删除文件失败！", ctx)
			return
		}
		return
	}
*/
