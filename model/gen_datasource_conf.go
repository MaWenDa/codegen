package model

import "time"

const (

	// GenDatasourceConfTableName 数据源表
	GenDatasourceConfTableName = "gen_datasource_conf"
)

// GenDatasourceConf 数据源结构
type GenDatasourceConf struct {
	ID         int64     `gorm:"column:id"          json:"id"`         // 主键ID
	Name       string    `gorm:"column:name"        json:"name"`       // 数据源名称
	URL        string    `gorm:"column:url"         json:"url"`        // jdbc-url
	Username   string    `gorm:"column:username"    json:"username"`   // 用户名
	Password   string    `gorm:"column:password"    json:"password"`   // 密码
	DelFlag    string    `gorm:"column:del_flag"    json:"delFlag"`    // 删除标记
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"` // 创建时间
	CreateBy   string    `gorm:"column:create_by"   json:"createBy"`   // 创建人
	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"` // 修改时间
	UpdateBy   string    `gorm:"column:update_by"   json:"updateBy"`   // 修改人
}
