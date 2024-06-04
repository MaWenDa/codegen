package tables

import (
	"codegen/api/define"
	"codegen/model"
	"codegen/pkg/db"
	"gorm.io/gorm"
)

func New() (*Dao, error) {
	return &Dao{db: db.Get("codegen")}, nil
}

type Dao struct {
	db *db.Group
}

// SelectCount 查询数量
func (dao *Dao) SelectCount(dataSourceTableListRequest *define.DataSourceTablePageRequest) int64 {

	count := int64(0)

	session := dao.StructureWhere(dataSourceTableListRequest)

	session.Count(&count)

	return count
}

// SelectPage 分页查询
func (dao *Dao) SelectPage(dataSourceTableListRequest *define.DataSourceTablePageRequest, current, size int) *[]define.DataSourceTableInfo {

	list := make([]define.DataSourceTableInfo, 0)

	total := dao.SelectCount(dataSourceTableListRequest)

	if total == 0 {
		return &list
	}

	// 计算分页参数
	offset := define.LimitParam(total, current, size)

	session := dao.StructureWhere(dataSourceTableListRequest)

	session.Order(define.DefaultOrder).
		Offset(offset).
		Limit(size).
		Find(&list)

	return &list
}

// SelectList 查询列表
func (dao *Dao) SelectList(dataSourceTableListRequest *define.DataSourceTablePageRequest) *[]define.DataSourceTableInfo {

	list := make([]define.DataSourceTableInfo, 0)

	session := dao.StructureWhere(dataSourceTableListRequest)

	session.Find(&list)

	return &list
}

// StructureWhere 构建Where
func (dao *Dao) StructureWhere(dataSourceTableListRequest *define.DataSourceTablePageRequest) *gorm.DB {

	session := dao.StructureSession().
		Where("table_schema = (select database())")

	if dataSourceTableListRequest.TableName != "" {
		session.Where("TABLE_NAME = ?", dataSourceTableListRequest.TableName)
	}

	return session
}

// StructureSession 构建Session
func (dao *Dao) StructureSession() *gorm.DB {

	session := dao.db.Slave().
		Table(model.TablesTableName).
		Model(&model.Tables{})

	return session
}

// StructureSessionByClient 构建Session
func (dao *Dao) StructureSessionByClient(client *db.Client) *gorm.DB {

	session := client.
		Table(model.TablesTableName).
		Model(&model.Tables{})

	return session
}
