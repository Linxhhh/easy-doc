package img_hash

import (
	"crypto/md5"
	"fmt"
	"mime/multipart"
	"io"
)

// 计算字节数据的 md5 值
func Md5(byteData []byte) string {
	hash := md5.New()
	hash.Write(byteData)
	hashByteData := hash.Sum(nil)
	return fmt.Sprintf("%x", hashByteData)
}

// 计算文件（图片）的 md5 值
func FileMd5(file multipart.File) string {
	hash := md5.New()
	io.Copy(hash, file)
	hashByteData := hash.Sum(nil)
	return fmt.Sprintf("%x", hashByteData)
}