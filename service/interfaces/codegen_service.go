package interfaces

import (
	"codegen/api/define"
	"codegen/pkg/server/api"
	"github.com/gin-gonic/gin"
)

type CodegenService interface {

	// TablePage 表数据-分页查询
	TablePage(ctx *gin.Context, req *define.DataSourceTablePageRequest) (*define.DataSourceTablePageResponse, api.Error)

	// GeneratorCode 生成代码
	GeneratorCode(ctx *gin.Context, req *define.GeneratorCodeRequest) api.Error
}
