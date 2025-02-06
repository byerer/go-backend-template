package config

import "go-backend-template/core/store/mysql"

type GlobalConfig struct {
	Port  string        `yaml:"Port"`
	Mysql mysql.OrmConf `yaml:"Mysql"`
}
