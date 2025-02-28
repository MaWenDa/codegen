package log

import (
	"context"
	"github.com/spf13/viper"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"path/filepath"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Config 配置文件
type Config struct {

	// 是否开发模式
	Development bool `json:"development" mapstructure:"development"`

	// 日志级别
	Level string `json:"level" mapstructure:"level"`

	// 日志输出路径
	Outputs []string `json:"outputs" mapstructure:"outputs"`

	// 日志最大体积 单位MB
	MaxSize int `json:"max_size" mapstructure:"max_size"`

	// 旧日志最长保存时间 0或不设置永久保存
	// MaxBackups 也会导致旧日志删除
	MaxAge int `json:"max_age" mapstructure:"max_age"`

	// 旧日志最多被打包数量 0或不设置永久保存
	// MaxAge 也会导致就日志删除
	MaxBackups int `json:"max_backups" mapstructure:"max_backups"`

	// 是否关闭sentry，默认是开启的
	DisableSentry bool `json:"disable_sentry" mapstructure:"disable_sentry"`
}

var (

	// 日志对象
	logger *zap.Logger

	// 配置文件
	config Config

	// 标准输出
	stdout = filepath.Base(os.Stdout.Name())

	// 标准错误输出
	stderr = filepath.Base(os.Stderr.Name())

	// 默认文件最大体积512MB 单位MB
	defaultMaxSize = 1 << 9

	// 允许输出的zap日志等级
	enableLevel zapcore.Level
)

func init() {
	logger, _ = zap.NewDevelopment()
}

func InitFromViper() {

	// 默认配置
	if err := viper.UnmarshalKey("log", &config); err != nil {
		panic(err)
	}

	// 日志最大体积
	maxSize := config.MaxSize
	if maxSize == 0 {
		maxSize = defaultMaxSize
	}

	// 日志等级标志转换
	switch config.Level {

	// 调试日志
	case "debug":
		enableLevel = zap.DebugLevel

	// 普通日志
	case "info":
		enableLevel = zap.InfoLevel

	// 警告日志
	case "warn":
		enableLevel = zap.WarnLevel

	// 错误日志
	case "error":
		enableLevel = zap.ErrorLevel
	}

	// 允许输出的日志等级
	levelEnabler := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= enableLevel
	})

	// 日志编码格式
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")

	// 内核集合
	cores := make([]zapcore.Core, 0)
	for _, outPath := range config.Outputs {
		switch outPath {

		// 标准输出
		case stdout:
			stdoutEncoder := zapcore.NewConsoleEncoder(encoderConfig)
			cores = append(cores, zapcore.NewCore(stdoutEncoder, zapcore.Lock(os.Stdout), levelEnabler))

		// 标准错误输出
		case stderr:
			stderrEncoder := zapcore.NewConsoleEncoder(encoderConfig)
			cores = append(cores, zapcore.NewCore(stderrEncoder, zapcore.Lock(os.Stdout), levelEnabler))

		// 默认文件输出
		default:

			// 文件写入对象
			lumberjackWriter := &lumberjack.Logger{
				Filename:   outPath,
				MaxSize:    maxSize,
				MaxAge:     config.MaxAge,
				MaxBackups: config.MaxBackups,
				LocalTime:  true,
				Compress:   true,
			}

			// 日志输出格式
			fileEncoder := zapcore.NewJSONEncoder(encoderConfig)
			cores = append(cores, zapcore.NewCore(fileEncoder, zapcore.AddSync(lumberjackWriter), levelEnabler))
		}
	}

	options := make([]zap.Option, 0)
	options = append(options, zap.AddCaller())

	// 创建日志实例
	logger = zap.New(zapcore.NewTee(cores...), options...)

}

// Logger 返回zap.Logger对象
func Logger() *zap.Logger {
	return logger
}

func Flush() {
	_ = logger.Sync()
}

func Sugar() *zap.SugaredLogger {
	return logger.Sugar()
}

func SugarContext(ctx context.Context) *zap.SugaredLogger {
	return logger.Sugar().With("request_id", ctx.Value("x-request-id"))
}
