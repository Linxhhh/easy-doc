package redis_service

import (
	"github.com/Linxhhh/easy-doc/global"
	"strconv"

	"github.com/go-redis/redis"
)

const (
	diggCount = "diggCount"
	readCount = "readCount"
)

func NewDocDigg() CountDB {
	return CountDB{
		Key: diggCount,
	}
}

func NewDocRead() CountDB {
	return CountDB{
		Key: readCount,
	}
}

/*
	     Key         Field       value
	浏览量/点赞量    文档ID        ...
*/

type CountDB struct {
	Key string
}

// 设置 Field，即文档 ID
func (c CountDB) SetById(docId uint) error {
	return c.Set(strconv.Itoa(int(docId)))
}

/* 
	设置、更新缓存数据，调用 + 1
*/
func (c CountDB) Set(field string) error {
	return c.SetNum(field, 1)
}

/* 
	设置、更新缓存数据，调用 + num
*/
func (c CountDB) SetNum(field string, num int) error {

	_, err := global.Redis.HIncrBy(c.Key, field, int64(num)).Result()
	if err == redis.Nil {
		_, err = global.Redis.HSet(c.Key, field, int64(num)).Result()
	}
	return err
}

// 返回该文档对应的缓存数据
func (c CountDB) GetById(docId uint) int {
	
	field := strconv.Itoa(int(docId))
	num, _ := global.Redis.HGet(c.Key, field).Int()
	return num
}

// 删除 Field 下的数据
func (c CountDB) Clear() {
	global.Redis.Del(c.Key)
}