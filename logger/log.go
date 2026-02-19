package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// 全局 Logger 实例
var Logger *zap.Logger

// InitLogger 初始化全局日志
func InitLogger() error {
	// 生产环境配置
	config := zap.NewProductionConfig()

	// ========== 可选优化配置 ==========

	// 设置日志级别（默认 Info）
	config.Level = zap.NewAtomicLevelAt(zap.InfoLevel)

	// 时间格式：ISO8601（默认已是）
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	// 输出到 stderr（默认）
	config.OutputPaths = []string{"stderr"}
	config.ErrorOutputPaths = []string{"stderr"}

	// 开发时可以打开（打印 caller）
	// config.EncoderConfig.EncodeCaller = zapcore.ShortCallerEncoder

	// =================================

	// 构建 logger
	logger, err := config.Build(
		zap.AddCaller(),      // 显示调用位置
		zap.AddCallerSkip(1), // 跳过封装层
	)
	if err != nil {
		return err
	}

	Logger = logger
	return nil
}

// Sync 刷新缓冲区（程序退出前调用）
func Sync() {
	if Logger != nil {
		_ = Logger.Sync()
	}
}
