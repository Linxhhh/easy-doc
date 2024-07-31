package global

import (
	"github.com/Linxhhh/easy-doc/config"

	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	Config *config.Config
	Log    *logrus.Logger
	DB     *gorm.DB
	Redis  *redis.Client
)
