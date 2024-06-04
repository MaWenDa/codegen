package datasource

import (
	"codegen/api/define"
	"codegen/model"
	"codegen/pkg/db"
	"gorm.io/gorm"
	"time"
)

func New() (*Dao, error) {
	return &Dao{db: db.Get("codegen")}, nil
}

type Dao struct {
	db *db.Group
}

// SelectById 根据id查询
func (dao *Dao) SelectById(id string) *model.GenDatasourceConf {

	genDatasourceConf := model.GenDatasourceConf{}

	dao.StructureSession().
		Where("del_flag = ?", "0").
		Where("id = ?", id).
		Find(&genDatasourceConf)

	return &genDatasourceConf
}

// SelectBatchIds 根据id集合查询
func (dao *Dao) SelectBatchIds(ids *[]string) *[]model.GenDatasourceConf {

	list := make([]model.GenDatasourceConf, 0)

	if len(*ids) > 0 {
		dao.StructureSession().
			Where("del_flag = ?", "0").
			Where("in (?)", ids).
			Find(&list)
	}

	return &list
}

// SelectList 查询列表
func (dao *Dao) SelectList(genDatasourceConf *model.GenDatasourceConf) *[]model.GenDatasourceConf {

	list := make([]model.GenDatasourceConf, 0)

	session := dao.StructureWhere(genDatasourceConf)

	session.Find(&list)

	return &list
}

// SelectCount 查询数量
func (dao *Dao) SelectCount(genDatasourceConf *model.GenDatasourceConf) int64 {

	count := int64(0)

	session := dao.StructureWhere(genDatasourceConf)

	session.Count(&count)

	return count
}

// SelectPage 分页查询
func (dao *Dao) SelectPage(genDatasourceConf *model.GenDatasourceConf, current, size int) *[]model.GenDatasourceConf {

	list := make([]model.GenDatasourceConf, 0)

	total := dao.SelectCount(genDatasourceConf)

	if total == 0 {
		return &list
	}

	// 计算分页参数
	offset := define.LimitParam(total, current, size)

	session := dao.StructureWhere(genDatasourceConf)

	session.Order(define.DefaultOrder).
		Offset(offset).
		Limit(size).
		Find(&list)

	return &list
}

// Insert 插入
func (dao *Dao) Insert(genDatasourceConf *model.GenDatasourceConf) bool {

	session := dao.StructureSessionByClient(dao.db.Master())

	genDatasourceConf.CreateTime = time.Now()
	genDatasourceConf.UpdateTime = time.Now()

	session.Create(genDatasourceConf)

	return true
}

// InsertBatchesByLength 批量插入
func (dao *Dao) InsertBatchesByLength(genDatasourceConfList []*model.GenDatasourceConf, length int) bool {

	session := dao.StructureSessionByClient(dao.db.Master())

	session.CreateInBatches(genDatasourceConfList, length).Set("create_time", time.Now())

	return true
}

// UpdateById 根据id更新
func (dao *Dao) UpdateById(genDatasourceConf *model.GenDatasourceConf) bool {

	if genDatasourceConf.ID == 0 {
		return dao.Insert(genDatasourceConf)
	}

	session := dao.StructureSessionByClient(dao.db.Master())

	genDatasourceConf.UpdateTime = time.Now()

	session.Where("id = ?", genDatasourceConf.ID).
		Updates(genDatasourceConf)

	return true
}

// DeleteById 根据id删除
func (dao *Dao) DeleteById(id string) bool {

	if id == "" {
		return false
	}

	session := dao.StructureSessionByClient(dao.db.Master())

	session.Where("id = ?", id).
		Delete(&model.GenDatasourceConf{})

	return true
}

// DeleteBatchIds 批量删除
func (dao *Dao) DeleteBatchIds(ids []string) bool {

	if len(ids) == 0 {
		return false
	}

	session := dao.StructureSessionByClient(dao.db.Master())

	session.Where("id in (?)", ids).
		Delete(&model.GenDatasourceConf{})

	return true
}

// StructureSession 构建Session
func (dao *Dao) StructureSession() *gorm.DB {

	session := dao.db.Slave().
		Table(model.GenDatasourceConfTableName).
		Model(&model.GenDatasourceConf{})

	return session
}

// StructureSessionByClient 构建Session
func (dao *Dao) StructureSessionByClient(client *db.Client) *gorm.DB {

	session := client.
		Table(model.GenDatasourceConfTableName).
		Model(&model.GenDatasourceConf{})

	return session
}

// StructureWhere 构建Where
func (dao *Dao) StructureWhere(genDatasourceConf *model.GenDatasourceConf) *gorm.DB {

	session := dao.StructureSession()

	if genDatasourceConf.ID != 0 {
		session.Where("id = ?", genDatasourceConf.ID)
	}

	if genDatasourceConf.Name != "" {
		session.Where("name like ?", "%"+genDatasourceConf.Name+"%")
	}

	return session
}
