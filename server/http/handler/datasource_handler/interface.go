package datasource_handler

import "github.com/gin-gonic/gin"

type DataSourceHandlerInterface interface {

	// SelectPage 分页查询
	SelectPage(ctx *gin.Context)

	// Insert 插入
	Insert(ctx *gin.Context)
}
