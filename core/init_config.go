package core

import (
	"github.com/Linxhhh/easy-doc/config"
	"os"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

const yamlPath = "settings.yaml"

func InitConfig() (c *config.Config) {
	byteData, err := os.ReadFile(yamlPath)
	if err != nil {
		logrus.Fatalln("read yaml err: ", err.Error())
	}

	c = new(config.Config)
	err = yaml.Unmarshal(byteData, c)
	if err != nil {
		logrus.Fatalln("解析 yaml err: ", err.Error())
	}

	return c
}