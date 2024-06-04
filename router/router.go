package router

import (
	dataSourceDao "codegen/dao/impl/datasource"
	"codegen/pkg/log"
	"codegen/pkg/server"
	"codegen/server/http/handler/codegen_handler"
	"codegen/server/http/handler/datasource_handler"
	"codegen/service/impl/codegen"
	"codegen/service/impl/datasource"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
)

func All() func(r *gin.Engine) {
	return func(r *gin.Engine) {

		// panic日志
		r.Use(ginzap.RecoveryWithZap(log.Sugar().Desugar(), true))

		prefixRouter := r.Group("/api/codegen")

		// 基础接口路由
		api(prefixRouter)
	}
}

func api(r *gin.RouterGroup) {

	// 数据源相关注入
	dsDao, err := dataSourceDao.New()
	if err != nil {
		log.Sugar().Error(err)
	}

	datasourceService, err := datasource.New(dsDao)
	if err != nil {
		log.Sugar().Error(err)
	}
	datasourceHandler, err := datasource_handler.GetDataSourceHandler(datasourceService)

	// 代码生成相关
	codegenService, err := codegen.New(dsDao)
	if err != nil {
		log.Sugar().Error(err)
	}
	codegenHandler, err := codegen_handler.GetCodegenHandler(codegenService)

	// 路由列表
	getHandlerMap := map[string]gin.HandlerFunc{
		"/datasource/page": datasourceHandler.SelectPage,
		"/codegen/page":    codegenHandler.TablePage,
	}

	postHandlerMap := map[string]gin.HandlerFunc{
		"/datasource/save":   datasourceHandler.Insert,
		"/codegen/generator": codegenHandler.GeneratorCode,
	}

	// 注册路由
	{
		options := r.Use(server.Cors())
		g := r.Use(server.Cors())
		for path, handler := range getHandlerMap {
			// 解决跨域
			options.OPTIONS(path)

			// 注册路由
			g.GET(path, handler)

		}

		for path, handler := range postHandlerMap {
			// 解决跨域
			options.OPTIONS(path)

			// 注册路由
			g.POST(path, handler)

		}
	}

}
