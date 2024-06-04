package config

import (
	"codegen/pkg/env"
	"codegen/pkg/utils/file"
	"github.com/spf13/viper"
)

// LoadConfigInFile 加载配置文件
func LoadConfigInFile(filename string) error {

	// 校验 filename 不为空，则按照默认规则加载配置文件
	if filename == "" {
		err := LoadConfig()
		if err != nil {
			return err
		}
	} else {
		viper.SetConfigFile(filename)
		if err := viper.ReadInConfig(); err != nil {
			return err
		}
	}

	return nil
}

// LoadConfig 加载配置文件
func LoadConfig() error {

	// 1. 读取配置文件 「app」
	if file.FilesExists([]string{"config/app.yml", "config/app.yaml", "config/app.json"}) {
		v := viper.New()
		v.AddConfigPath("config")
		v.SetConfigName("app")
		if err := v.ReadInConfig(); err != nil {
			return err
		}

		// 加载的所有配置合并到全局的配置中
		err := viper.MergeConfigMap(v.AllSettings())
		if err != nil {
			return err
		}
	}

	// 获取当前运行环境
	env := env.GetEnvironment()

	// 2. 读取配置文件 「env」 ，会覆盖 app 里的配置
	if file.FilesExists([]string{"config/" + env + ".yml", "config/" + env + ".yaml", "config/" + env + ".json"}) {
		v := viper.New()
		v.AddConfigPath("config")
		v.SetConfigName(env)
		if err := v.ReadInConfig(); err != nil {
			return err
		}

		// 加载的所有配置合并到全局的配置中
		err := viper.MergeConfigMap(v.AllSettings())
		if err != nil {
			return err
		}
	}

	return nil
}

func Debug() bool {
	return viper.GetBool("debug")
}
