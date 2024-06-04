package interfaces

import (
	"codegen/model"
	"codegen/pkg/db"
	"gorm.io/gorm"
)

type DatasourceDao interface {

	// SelectById 根据id获取
	SelectById(id string) *model.GenDatasourceConf

	// SelectBatchIds 根据id集合查询
	SelectBatchIds(ids *[]string) *[]model.GenDatasourceConf

	// SelectList 查询列表
	SelectList(genDatasourceConf *model.GenDatasourceConf) *[]model.GenDatasourceConf

	// SelectCount 查询数量
	SelectCount(genDatasourceConf *model.GenDatasourceConf) int64

	// SelectPage 分页查询
	SelectPage(genDatasourceConf *model.GenDatasourceConf, current, size int) *[]model.GenDatasourceConf

	// Insert 插入
	Insert(genDatasourceConf *model.GenDatasourceConf) bool

	// InsertBatchesByLength 批量插入
	InsertBatchesByLength(genDatasourceConfList []*model.GenDatasourceConf, length int) bool

	// UpdateById 修改
	UpdateById(genDatasourceConf *model.GenDatasourceConf) bool

	// DeleteById 根据id删除
	DeleteById(id string) bool

	// DeleteBatchIds 根据id集合删除
	DeleteBatchIds(ids []string) bool

	// StructureSession 构建Session
	StructureSession() *gorm.DB

	// StructureWhere 构建Where
	StructureWhere(genDatasourceConf *model.GenDatasourceConf) *gorm.DB

	// StructureSessionByClient 构建Session
	StructureSessionByClient(client *db.Client) *gorm.DB
}
