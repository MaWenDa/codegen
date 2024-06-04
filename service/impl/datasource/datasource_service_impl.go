package datasource

import (
	"codegen/api/define"
	errs "codegen/api/error"
	dao "codegen/dao/interfaces"
	"codegen/model"
	"codegen/pkg/db"
	"codegen/pkg/server/api"
)

func New(dd dao.DatasourceDao) (*Service, error) {
	return &Service{
		datasourceDao: dd,
	}, nil
}

type Service struct {
	datasourceDao dao.DatasourceDao
}

// SelectPage 分页查询
func (service *Service) SelectPage(pageReq *define.BasePageRequest, genDatasourceConf *model.GenDatasourceConf) (*define.BasePageResponse, api.Error) {

	// 校验分页参数
	pageReq.Size, pageReq.Current = define.VerifyLimit(pageReq.Size, pageReq.Current)

	// 初始化响应数据
	pageResp := &define.BasePageResponse{List: make([]model.GenDatasourceConf, 0)}
	pageResp.Size = pageReq.Size
	pageResp.Current = pageReq.Current

	// 查询总数
	pageResp.Total = service.datasourceDao.SelectCount(genDatasourceConf)
	if pageResp.Total == 0 {
		return pageResp, nil
	}

	// 分页查询
	pageResp.List = service.datasourceDao.SelectPage(genDatasourceConf, pageReq.Current, pageReq.Size)

	return pageResp, nil
}

// SelectById 根据id获取
func (service *Service) SelectById(id string) (*model.GenDatasourceConf, api.Error) {
	return service.datasourceDao.SelectById(id), nil
}

// SelectList 查询列表
func (service *Service) SelectList(genDatasourceConf *model.GenDatasourceConf) (*[]model.GenDatasourceConf, api.Error) {
	return service.datasourceDao.SelectList(genDatasourceConf), nil
}

// Insert 插入
func (service *Service) Insert(genDatasourceConf *model.GenDatasourceConf) (bool, api.Error) {
	if !db.TestConnMysql(genDatasourceConf.Username, genDatasourceConf.Password, genDatasourceConf.URL) {
		return false, errs.DatasourceConnFailed
	}
	return service.datasourceDao.Insert(genDatasourceConf), nil
}

// UpdateById 修改
func (service *Service) UpdateById(genDatasourceConf *model.GenDatasourceConf) (bool, api.Error) {
	return service.datasourceDao.UpdateById(genDatasourceConf), nil
}

// DeleteById 根据id删除
func (service *Service) DeleteById(id string) (bool, api.Error) {
	return service.datasourceDao.DeleteById(id), nil
}

// DeleteBatchIds 根据id集合删除
func (service *Service) DeleteBatchIds(ids []string) (bool, api.Error) {
	return service.datasourceDao.DeleteBatchIds(ids), nil
}
