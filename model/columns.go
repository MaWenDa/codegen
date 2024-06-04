package model

const (

	// ColumnsTableName 字段表
	ColumnsTableName = "information_schema.columns"
)

// Columns 字段元数据信息
type Columns struct {
	ColumnName    string `gorm:"column:COLUMN_NAME" json:"columnName"`       // 字段名
	DataType      string `gorm:"column:DATA_TYPE" json:"dataType"`           // 数据类型
	ColumnComment string `gorm:"column:COLUMN_COMMENT" json:"columnComment"` // 字段注释
	ColumnKey     string `gorm:"column:COLUMN_KEY" json:"columnKey"`         // 字段键值
	Extra         string `gorm:"column:EXTRA" json:"extra"`                  // 额外信息
	IsNullable    string `gorm:"column:IS_NULLABLE" json:"isNullable"`       // 是否可为空
	ColumnType    string `gorm:"column:COLUMN_TYPE" json:"columnType"`       // 字段类型
}
