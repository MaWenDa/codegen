package colums

import (
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

// StructureSession 构建Session
func (dao *Dao) StructureSession() *gorm.DB {

	session := dao.db.Slave().
		Table(model.ColumnsTableName).
		Model(&model.Columns{})

	return session
}

// StructureSessionByClient 构建Session
func (dao *Dao) StructureSessionByClient(client *db.Client) *gorm.DB {

	session := client.
		Table(model.ColumnsTableName).
		Model(&model.Columns{})

	return session
}
