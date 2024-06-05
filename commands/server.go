package commands

import (
	"codegen/pkg/graceful"
	"codegen/pkg/server"
	"codegen/router"
	"github.com/urfave/cli/v2"
)

// Serve 启动HTTP服务
func Serve(c *cli.Context) error {

	// 运行HTTP服务
	graceful.Start(server.NewHttp(server.Addr(":8060"), server.Router(router.All())))

	graceful.Wait()

	return nil
}
