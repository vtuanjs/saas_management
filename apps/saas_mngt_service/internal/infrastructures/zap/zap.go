package zap

import (
	"context"
	"os"

	system "github.com/vtuanjs/saas_management/apps/saas_mngt_service/internal/core"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type loggerZap struct {
	*zap.Logger
}

func (l *loggerZap) Debug(_ context.Context, msg string, data ...any) {
	l.Logger.Debug(msg, zap.Any("data", data))
}

func (l *loggerZap) Info(_ context.Context, msg string, data ...any) {
	l.Logger.Info(msg, zap.Any("data", data))
}

func (l *loggerZap) Warn(_ context.Context, msg string, data ...any) {
	l.Logger.Warn(msg, zap.Any("data", data))
}

func (l *loggerZap) Error(_ context.Context, msg string, err error) {
	l.Logger.Error(msg, zap.Error(err))
}

func (l *loggerZap) Fatal(_ context.Context, msg string, err error) {
	l.Logger.Fatal(msg, zap.Error(err))
}

func NewLogger(cfg *system.Config) system.Logger {
	logCfg := cfg.Logger
	var level zapcore.Level
	switch logCfg.Level {
	case "debug":
		level = zapcore.DebugLevel
	case "info":
		level = zapcore.InfoLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	default:
		level = zapcore.InfoLevel
	}

	encoder := getEncoderLog()
	hook := lumberjack.Logger{
		Filename:   logCfg.Filename,
		MaxSize:    logCfg.MaxSize,
		MaxBackups: logCfg.MaxBackups,
		MaxAge:     logCfg.MaxAge,
		Compress:   logCfg.Compress,
	}
	core := zapcore.NewCore(
		encoder,
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)),
		level,
	)

	return &loggerZap{zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))}
}

func getEncoderLog() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}
