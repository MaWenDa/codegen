package string

import (
	"codegen/pkg/utils/file"
	"os"
	"strings"
	"text/template"
)

// ConvertToCamelCase 将字符串转换为驼峰命名法
func ConvertToCamelCase(input string) string {
	// 将输入字符串按下划线分割成单词
	words := strings.Split(input, "_")

	// 遍历单词列表，将每个单词首字母大写
	var result string
	for _, word := range words {
		firstLetter := strings.ToUpper(string(word[0]))
		rest := word[1:]
		result += firstLetter + rest
	}

	return result
}

// ConvertToHump 将字符串转换为驼峰命名法
func ConvertToHump(input string) string {
	// 将输入字符串按下划线分割成单词
	words := strings.Split(input, "_")

	// 遍历单词列表，将除了第一个单词外的其他单词的首字母大写
	var result string
	for i, word := range words {
		if i == 0 {
			result += word
		} else {
			firstLetter := strings.ToUpper(string(word[0]))
			rest := word[1:]
			result += firstLetter + rest
		}
	}

	return result
}

// MapDBTypeToGo 将数据库类型映射到Go类型
func MapDBTypeToGo(dbType string) string {
	switch dbType {
	case "int", "tinyint", "smallint", "mediumint":
		return "int"
	case "bigint":
		return "int64"
	case "float":
		return "float32"
	case "double", "decimal":
		return "float64"
	case "char", "varchar", "text":
		return "string"
	case "date", "time", "datetime", "timestamp":
		return "time.Time"
	case "boolean", "bool":
		return "bool"
	default:
		return "interface{}" // 默认情况下返回空接口类型
	}
}

// GeneratorCodeTemplateFilling 模板数据填充
func GeneratorCodeTemplateFilling(data interface{}, filePath, templatePath string) bool {
	// 解析模板文件
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return false
	}

	// 创建输出文件
	_ = file.CreateFile(filePath)

	outFile, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}

	defer outFile.Close()

	// 应用模板并生成代码到输出文件
	err = tmpl.Execute(outFile, data)
	if err != nil {
		return false
	}

	return true
}
