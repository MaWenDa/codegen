package db

import "github.com/spf13/viper"

// ClientConfig 数据库配置
type ClientConfig struct {
	Dsn string `json:"dsn"`
}

// Config 数据库链接配置
type Config struct {

	// 主库
	Master ClientConfig `json:"master"`

	// 从库集合
	Slaves []ClientConfig `json:"slaves"`
}

// InitFromViper 从viper中初始化数据库配置
func InitFromViper() {

	// 解析数据库配置
	dbCfg := make(map[string]Config)
	err := viper.UnmarshalKey("db.databases", &dbCfg)
	if err != nil {
		panic(err)
	}

	// 数据库链接组添加到GroupManager
	for name, cfg := range dbCfg {

		// 实例化数据库链接组
		group, err := NewGroup(name, cfg)
		if err != nil {
			panic(err)
		}

		// 添加到GroupManager
		groupManager.Add(name, group)
	}
}
