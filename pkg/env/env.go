package env

import "os"

// GetEnvironment 获取当前运行环境，默认为 develop
func GetEnvironment() string {

	// 获取 ENV 环境变量
	env := os.Getenv("ENV")

	switch env {
	case "prod", "production":
		return "production"
	case "test":
		return "test"
	case "dev", "develop":
		return "develop"
	default:
		return "develop"
	}
}

// GetShortEnvironment 获取当前运行环境简称，默认为 dev
func GetShortEnvironment() string {

	// 获取 ENV 环境变量
	env := GetEnvironment()

	switch env {
	case "prod", "production":
		return "prod"
	case "test":
		return "test"
	case "dev", "develop":
		return "dev"
	default:
		return "dev"
	}
}

// IsProduction 是否生产环境
func IsProduction() bool {
	return GetEnvironment() == "production"
}

// IsTest 是否测试环境
func IsTest() bool {
	return GetEnvironment() == "test"
}
