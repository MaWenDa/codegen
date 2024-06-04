package codegen_handler

import "github.com/gin-gonic/gin"

type CodegenHandlerInterface interface {

	// TablePage 分页查询
	TablePage(ctx *gin.Context)

	// GeneratorCode 代码生成
	GeneratorCode(ctx *gin.Context)
}
