package model

import "time"

const (

	// TablesTableName 表元数据信息表
	TablesTableName = "information_schema.tables"
)

// Tables 表元数据信息
type Tables struct {
	TableName      string    `gorm:"column:TABLE_NAME"  json:"tableName"`           // 表名
	TableComment   string    `gorm:"column:TABLE_COMMENT"  json:"tableComment"`     // 表注释
	Engine         string    `gorm:"column:ENGINE"  json:"engine"`                  // 引擎
	TableCollation string    `gorm:"column:TABLE_COLLATION"  json:"tableCollation"` // 表编码
	CreateTime     time.Time `gorm:"column:CREATE_TIME"  json:"createTime"`         // 创建时间
}
