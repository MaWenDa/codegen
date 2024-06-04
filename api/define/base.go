package define

import "math"

const (
	// DefaultOrder 默认排序
	DefaultOrder = "create_time DESC"
)

// BasePageResponse 数据源列表响应
type BasePageResponse struct {
	Total   int64       `json:"total"`   // 总行数
	Current int         `json:"current"` // 当前页
	Size    int         `json:"size"`    // 展示行数
	List    interface{} `json:"list"`    // 数据源列表
}

// BasePageRequest 数据源列表请求
type BasePageRequest struct {
	Total   int64 `json:"total"   form:"total"`   // 总行数
	Current int   `json:"current" form:"current"` // 当前页
	Size    int   `json:"size"    form:"size"`    // 展示行数
}

// LimitParam 计算分页参数
func LimitParam(total int64, current, size int) (offset int) {

	// 计算最大页数
	maxPage := int(math.Ceil(float64(total) / float64(size)))

	// 当前页数大于最大页数
	if current > maxPage {
		current = maxPage
	}
	// 计算偏移量
	offset = (current - 1) * size

	return offset
}

// VerifyLimit 校验分页参数
func VerifyLimit(size, current int) (int, int) {

	// 每页最少1条数据
	if size < 1 {
		size = 1
	}

	// 每页最多20条数据
	if size > 100 {
		size = 100
	}

	// 当前页
	if current < 1 {
		current = 1
	}

	return size, current
}
