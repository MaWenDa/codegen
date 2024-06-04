package define

import (
	"time"
)

// DataSource 数据源信息
type DataSource struct {
	ID         int64     `json:"id" gorm:"column:id"`                  // 主键ID
	Name       string    `json:"name" gorm:"column:name"`              // 数据源名称
	URL        string    `json:"url" gorm:"column:url"`                // jdbc-url
	Username   string    `json:"username" gorm:"column:username"`      // 用户名
	CreateTime time.Time `json:"createTime" gorm:"column:create_time"` // 创建时间
	CreateBy   string    `json:"createBy" gorm:"column:create_by"`     // 创建人
}

// DataSourceListResponse 数据源列表响应
type DataSourceListResponse struct {
	Total   int64        `json:"total"`   // 总行数
	Current int          `json:"current"` // 当前页
	Size    int          `json:"size"`    // 展示行数
	List    []DataSource `json:"list"`    // 数据源列表
}

// DataSourceListRequest 数据源列表请求
type DataSourceListRequest struct {
	Total   int64  `json:"total"   form:"total"`   // 总行数
	Current int    `json:"current" form:"current"` // 当前页
	Size    int    `json:"size"    form:"size"`    // 展示行数
	Name    string `json:"name"    form:"name"`    // 数据源名称
}

// DataSourceSaveRequest 数据源保存请求
type DataSourceSaveRequest struct {
	Name     string `json:"name"     form:"name"     binding:"required"` // 数据源名称
	URL      string `json:"url"      form:"url"      binding:"required"` // jdbc-url
	Username string `json:"username" form:"username" binding:"required"` // 用户名
	Password string `json:"password" form:"password" binding:"required"` // 密码
}
