package db

import (
	"fmt"
	"regexp"
)

// ExtractDatabaseName 从数据源名称中提取数据库名称
func ExtractDatabaseName(dataSourceName string) (string, error) {
	// 定义正则表达式
	re := regexp.MustCompile(`@tcp\([^:]+:[^)]+\)/([^?]+)`)

	// 在连接字符串中查找匹配的子串
	match := re.FindStringSubmatch(dataSourceName)
	if len(match) != 2 {
		return "", fmt.Errorf("无法从连接字符串中提取数据库名称")
	}

	// 返回数据库名称
	return match[1], nil
}
