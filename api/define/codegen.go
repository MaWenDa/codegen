package define

import "time"

const (
	OUT_PATH              = "/Users/mawenda/Documents/tmp/out/"
	MODEL_TEMPLATE        = "config/template/model.tmpl"
	DAO_TEMPLATE          = "config/template/dao.tmpl"
	DAO_IMPL_TEMPLATE     = "config/template/dao_impl.tmpl"
	SERVICE_TEMPLATE      = "config/template/service.tmpl"
	SERVICE_IMPL_TEMPLATE = "config/template/service_impl.tmpl"
	HANDLER_TEMPLATE      = "config/template/handler.tmpl"
	HANDLER_IMPL_TEMPLATE = "config/template/handler_impl.tmpl"
)

// TableModel 表属性信息
type TableModel struct {
	TableName     string        `gorm:"column:TABLE_NAME"`    // 表名
	TableComment  string        `gorm:"column:TABLE_COMMENT"` // 表注释
	ModelName     string        `gorm:"-"`                    // 结构体名称
	ProjectName   string        `gorm:"-"`                    // 项目名称
	DatabaseName  string        `gorm:"-"`                    // 数据库名称
	ModelNameHump string        `gorm:"-"`                    // 结构体驼峰命名
	ColumnModels  []ColumnModel `gorm:"-"`                    // 列信息
}

// ColumnModel 列属性信息
type ColumnModel struct {
	ColumnName      string `gorm:"column:COLUMN_NAME"`    // 列名
	ColumnComment   string `gorm:"column:COLUMN_COMMENT"` // 列注释
	DataType        string `gorm:"column:DATA_TYPE"`      // 数据类型
	ColumnKey       string `gorm:"column:COLUMN_KEY"`     // 字段键值
	ColumnModelName string `gorm:"-"`                     // 结构体字段名称
	ColumnType      string `gorm:"-"`                     // 字段类型
	ColumnNameHump  string `gorm:"-"`                     // 列名驼峰命名
	ModelNameHump   string `gorm:"-"`                     // 结构体驼峰命名
	CheckEmpty      string `gorm:"-"`                     // 判断空值
}

// DataSourceTablePageRequest 数据源表列表请求
type DataSourceTablePageRequest struct {
	Total        int64  `json:"total"        form:"total"`                           // 总行数
	Current      int    `json:"current"      form:"current"`                         // 当前页
	Size         int    `json:"size"         form:"size"`                            // 展示行数
	DataSourceId string `json:"dataSourceId" form:"dataSourceId" binding:"required"` // 数据源ID
	TableName    string `json:"tableName"    form:"tableName"`                       // 表名
}

// DataSourceTablePageResponse 数据源表列表响应
type DataSourceTablePageResponse struct {
	Total   int64                 `json:"total"`   // 总行数
	Current int                   `json:"current"` // 当前页
	Size    int                   `json:"size"`    // 展示行数
	List    []DataSourceTableInfo `json:"list"`    // 数据源列表
}

// DataSourceTableInfo 表信息
type DataSourceTableInfo struct {
	TableName      string    `gorm:"column:TABLE_NAME"  json:"tableName"`           // 表名
	TableComment   string    `gorm:"column:TABLE_COMMENT"  json:"tableComment"`     // 表注释
	Engine         string    `gorm:"column:ENGINE"  json:"engine"`                  // 引擎
	TableCollation string    `gorm:"column:TABLE_COLLATION"  json:"tableCollation"` // 表编码
	CreateTime     time.Time `gorm:"column:CREATE_TIME"  json:"createTime"`         // 创建时间
}

// DataSourceTableColumnInfo 表字段信息
type DataSourceTableColumnInfo struct {
	ColumnName    string `gorm:"column:COLUMN_NAME" json:"columnName"`       // 字段名
	DataType      string `gorm:"column:DATA_TYPE" json:"dataType"`           // 数据类型
	ColumnComment string `gorm:"column:COLUMN_COMMENT" json:"columnComment"` // 字段注释
	ColumnKey     string `gorm:"column:COLUMN_KEY" json:"columnKey"`         // 字段键值
	Extra         string `gorm:"column:EXTRA" json:"extra"`                  // 额外信息
	IsNullable    string `gorm:"column:IS_NULLABLE" json:"isNullable"`       // 是否可为空
	ColumnType    string `gorm:"column:COLUMN_TYPE" json:"columnType"`       // 字段类型
}

// GeneratorCodeRequest 生成代码请求
type GeneratorCodeRequest struct {
	DataSourceId string `json:"dataSourceId" binding:"required"` // 数据源ID
	TableName    string `json:"tableName" binding:"required"`    // 表名
	ProjectName  string `json:"projectName" binding:"required"`  // 项目名称
}
