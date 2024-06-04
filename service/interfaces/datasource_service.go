package interfaces

import (
	"codegen/api/define"
	"codegen/model"
	"codegen/pkg/server/api"
)

// DataSourceService 数据源接口声明
type DataSourceService interface {

	// SelectPage 分页查询
	SelectPage(pageReq *define.BasePageRequest, genDatasourceConf *model.GenDatasourceConf) (*define.BasePageResponse, api.Error)

	// SelectById 根据id获取
	SelectById(id string) (*model.GenDatasourceConf, api.Error)

	// SelectList 查询列表
	SelectList(genDatasourceConf *model.GenDatasourceConf) (*[]model.GenDatasourceConf, api.Error)

	// Insert 插入
	Insert(genDatasourceConf *model.GenDatasourceConf) (bool, api.Error)

	// UpdateById 修改
	UpdateById(genDatasourceConf *model.GenDatasourceConf) (bool, api.Error)

	// DeleteById 根据id删除
	DeleteById(id string) (bool, api.Error)

	// DeleteBatchIds 根据id集合删除
	DeleteBatchIds(ids []string) (bool, api.Error)
}
