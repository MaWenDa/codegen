package codegen_handler

import (
	"codegen/api/define"
	errs "codegen/api/error"
	"codegen/pkg/log"
	"codegen/service/interfaces"
	"github.com/gin-gonic/gin"
)

// GetCodegenHandler 依赖注入
func GetCodegenHandler(ss interfaces.CodegenService) (CodegenHandlerInterface, error) {
	return &CodegenHandler{
		service: ss,
	}, nil
}

type CodegenHandler struct {
	service interfaces.CodegenService
}

// TablePage 分页查询数据表
func (codegenHandler *CodegenHandler) TablePage(ctx *gin.Context) {

	// 获取请求参数
	req := define.DataSourceTablePageRequest{}
	if err := ctx.ShouldBind(&req); err != nil {
		log.SugarContext(ctx).Errorw("CodegenHandler TablePage ShouldBind", "err", err)
		ctx.JSON(errs.ErrResp(errs.ParamErr))
		return
	}

	// 执行查询
	resp, err := codegenHandler.service.TablePage(ctx, &req)
	if err != nil {
		ctx.JSON(errs.ErrResp(err))
		return
	}

	ctx.PureJSON(errs.SucResp(resp))
	return
}

// GeneratorCode 代码生成
func (codegenHandler *CodegenHandler) GeneratorCode(ctx *gin.Context) {
	// 获取请求参数
	req := define.GeneratorCodeRequest{}
	if err := ctx.ShouldBind(&req); err != nil {
		log.SugarContext(ctx).Errorw("CodegenHandler GeneratorCode ShouldBind", "err", err)
		ctx.JSON(errs.ErrResp(errs.ParamErr))
		return
	}

	err := codegenHandler.service.GeneratorCode(ctx, &req)
	if err != nil {
		ctx.JSON(errs.ErrResp(err))
		return
	}
	ctx.PureJSON(errs.SucResp(true))
	return
}
