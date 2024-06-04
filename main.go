package main

import (
	"codegen/commands"
	"codegen/pkg/config"
	"codegen/pkg/db"
	"codegen/pkg/log"
	"github.com/urfave/cli/v2"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Action = commands.Serve
	app.Before = initConfig
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:     "config",
			Value:    "config",
			Usage:    "specify the location of the configuration file",
			Required: false,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Sugar().Fatal(err)
	}
	log.Flush()
}

// 初始化配置
func initConfig(*cli.Context) error {

	// 读取配置
	if err := config.LoadConfig(); err != nil {
		return err
	}

	// 加载日志配置
	log.InitFromViper()

	// 加载数据库配置
	db.InitFromViper()

	return nil
}
