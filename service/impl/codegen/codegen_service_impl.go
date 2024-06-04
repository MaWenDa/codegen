package codegen

import (
	"codegen/api/define"
	errs "codegen/api/error"
	dao "codegen/dao/interfaces"
	"codegen/model"
	"codegen/pkg/db"
	"codegen/pkg/log"
	"codegen/pkg/server/api"
	dbUtil "codegen/pkg/utils/db"
	generatorUtils "codegen/pkg/utils/generator"
	"github.com/gin-gonic/gin"
)

func New(dd dao.DatasourceDao) (*Service, error) {
	return &Service{
		datasourceDao: dd,
	}, nil
}

type Service struct {
	datasourceDao dao.DatasourceDao
}

// TablePage 表列表
func (service *Service) TablePage(ctx *gin.Context, req *define.DataSourceTablePageRequest) (*define.DataSourceTablePageResponse, api.Error) {

	// 校验分页参数
	req.Size, req.Current = define.VerifyLimit(req.Size, req.Current)

	response := &define.DataSourceTablePageResponse{}
	response.Size = req.Size
	response.Current = req.Current

	// 获取数据源信息
	datasource := service.datasourceDao.SelectById(req.DataSourceId)

	// 数据源信息为空时，返回提示
	if datasource.ID == 0 {
		return response, errs.RecordNotFound
	}

	// 根据数据源信息连接数据库
	db := db.ConnMysql(datasource.Username, datasource.Password, datasource.URL)

	// 构造查询
	session := db.Table(model.TablesTableName).Model(&model.Tables{}).Where("table_schema = (select database())")

	// 构建查询条件
	if req.TableName != "" {
		session.Where("table_name like ?", "%"+req.TableName+"%")
	}

	// 查询总数
	if err := session.Count(&response.Total).Error; err != nil {
		log.SugarContext(ctx).Errorw("TablesService TablePage Count", "TableName", req.TableName, "err", err)
		return response, errs.Failed
	}

	// 总数为0，直接返回
	if response.Total == 0 {
		return response, nil
	}

	// 计算分页参数
	offset := define.LimitParam(response.Total, req.Current, req.Size)

	// 执行查询
	err := session.
		Order(define.DefaultOrder).
		Offset(offset).
		Limit(req.Size).
		Find(&response.List).
		Error

	// 处理查询错误
	if err != nil {
		log.SugarContext(ctx).Errorw("TablesService TablePage Find", "req", req, "err", err)
		return response, errs.Failed
	}

	return response, nil
}

// GeneratorCode 生成代码
func (service *Service) GeneratorCode(ctx *gin.Context, req *define.GeneratorCodeRequest) api.Error {

	datasource := service.datasourceDao.SelectById(req.DataSourceId)

	// 数据源信息为空时，返回提示
	if datasource.ID == 0 {
		return errs.RecordNotFound
	}

	db := db.ConnMysql(datasource.Username, datasource.Password, datasource.URL)

	tableModels := make([]define.TableModel, 0)
	// 构造查询
	db.Table(model.TablesTableName).
		Model(&model.Tables{}).
		Select("table_name, table_comment, engine, table_collation, create_time").
		Where("table_schema = (select database())").
		Where("table_name in (?) ", req.TableName).
		Find(&tableModels)

	if len(tableModels) == 0 {
		return errs.RecordNotFound
	}

	// 解析数据库名称
	dataBaseName, err := dbUtil.ExtractDatabaseName(datasource.URL)
	if err != nil {
		log.SugarContext(ctx).Errorw("ExtractDatabaseName", "err", err)
		return errs.Failed
	}

	for _, tableModel := range tableModels {

		// 封装表元数据信息
		tableModel.ModelName = generatorUtils.ConvertToCamelCase(tableModel.TableName)
		tableModel.ModelNameHump = generatorUtils.ConvertToHump(tableModel.TableName)
		tableModel.ProjectName = req.ProjectName
		tableModel.DatabaseName = dataBaseName

		// 获取表字段信息
		tableModel.ColumnModels = make([]define.ColumnModel, 0)
		db.Table(model.ColumnsTableName).
			Model(&model.Columns{}).
			Select("column_name, data_type, column_comment, data_type").
			Order("ORDINAL_POSITION ASC").
			Where("table_schema = (select database())").
			Where("table_name = ?", tableModel.TableName).
			Find(&tableModel.ColumnModels)
		// 对表字段信息封装
		for i := range tableModel.ColumnModels {
			tableModel.ColumnModels[i].ColumnModelName = generatorUtils.ConvertToCamelCase(tableModel.ColumnModels[i].ColumnName)
			tableModel.ColumnModels[i].ColumnType = generatorUtils.MapDBTypeToGo(tableModel.ColumnModels[i].DataType)
			tableModel.ColumnModels[i].ColumnNameHump = generatorUtils.ConvertToHump(tableModel.ColumnModels[i].ColumnName)
			tableModel.ColumnModels[i].ModelNameHump = tableModel.ModelNameHump
		}

		// 生成代码
		generatorCode(&tableModel)

	}

	return nil
}

// generatorCode 生成代码
func generatorCode(tableModel *define.TableModel) bool {

	// 文件输出路径
	outPath := define.OUT_PATH + tableModel.TableName + "/"

	var fileName = tableModel.TableName + ".go"

	// 生成model
	generatorUtils.GeneratorCodeTemplateFilling(tableModel, outPath+"model/"+fileName, define.MODEL_TEMPLATE)

	// 生成dao
	generatorUtils.GeneratorCodeTemplateFilling(tableModel, outPath+"dao/interfaces/"+fileName, define.DAO_TEMPLATE)
	generatorUtils.GeneratorCodeTemplateFilling(tableModel, outPath+"dao/impl/"+tableModel.TableName+"/"+fileName, define.DAO_IMPL_TEMPLATE)

	// 生成service
	generatorUtils.GeneratorCodeTemplateFilling(tableModel, outPath+"service/interfaces/"+fileName, define.SERVICE_TEMPLATE)
	generatorUtils.GeneratorCodeTemplateFilling(tableModel, outPath+"service/impl/"+tableModel.TableName+"/"+fileName, define.SERVICE_IMPL_TEMPLATE)

	// 生成handler
	generatorUtils.GeneratorCodeTemplateFilling(tableModel, outPath+"server/http/handler/"+tableModel.TableName+"/interfaces"+".go", define.HANDLER_TEMPLATE)
	generatorUtils.GeneratorCodeTemplateFilling(tableModel, outPath+"server/http/handler/"+tableModel.TableName+"/"+fileName, define.HANDLER_IMPL_TEMPLATE)

	return true
}
