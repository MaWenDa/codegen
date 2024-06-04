package datasource_handler

import (
	"codegen/api/define"
	errs "codegen/api/error"
	"codegen/model"
	"codegen/pkg/log"
	"codegen/service/interfaces"
	"github.com/gin-gonic/gin"
)

// GetDataSourceHandler 依赖注入
func GetDataSourceHandler(ss interfaces.DataSourceService) (DataSourceHandlerInterface, error) {
	return &DataSourceHandler{
		service: ss,
	}, nil
}

type DataSourceHandler struct {
	service interfaces.DataSourceService
}

// SelectPage 分页查询
func (dataSourceHandler *DataSourceHandler) SelectPage(ctx *gin.Context) {

	// 获取分页参数
	basePageRequest := define.BasePageRequest{}
	if err := ctx.ShouldBind(&basePageRequest); err != nil {
		log.SugarContext(ctx).Errorw("DataSourceHandler SelectPage BasePageRequest ShouldBind", "err", err)
		ctx.JSON(errs.ErrResp(errs.ParamErr))
		return
	}

	// 获取筛选参数
	genDatasourceConf := model.GenDatasourceConf{}
	if err := ctx.ShouldBind(&genDatasourceConf); err != nil {
		log.SugarContext(ctx).Errorw("DataSourceHandler SelectPage GenDatasourceConf ShouldBind", "err", err)
		ctx.JSON(errs.ErrResp(errs.ParamErr))
		return
	}

	// 执行查询
	resp, err := dataSourceHandler.service.SelectPage(&basePageRequest, &genDatasourceConf)
	if err != nil {
		ctx.JSON(errs.ErrResp(err))
		return
	}

	ctx.PureJSON(errs.SucResp(resp))
	return
}

// Insert 插入
func (dataSourceHandler *DataSourceHandler) Insert(ctx *gin.Context) {

	// 获取请求参数
	genDatasourceConf := model.GenDatasourceConf{}
	if err := ctx.ShouldBind(&genDatasourceConf); err != nil {
		log.SugarContext(ctx).Errorw("DataSourceHandler Insert ShouldBind", "err", err)
		ctx.JSON(errs.ErrResp(errs.ParamErr))
		return
	}

	// 执行插入
	resp, err := dataSourceHandler.service.Insert(&genDatasourceConf)
	if err != nil {
		ctx.JSON(errs.ErrResp(err))
		return
	}

	ctx.PureJSON(errs.SucResp(resp))
	return
}
