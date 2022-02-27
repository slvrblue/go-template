package logger

import (
	"errors"
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/blattaria7/go-template/config"
)

const (
	// TraceLevel is a trace logger level.
	TraceLevel = "trace"

	// DebugLevel is a debug logger level.
	DebugLevel = "debug"

	// InfoLevel is an info logger level.
	InfoLevel = "info"

	// WarningLevel is a warning logger level.
	WarningLevel = "warning"

	// ErrorLevel is an error logger level.
	ErrorLevel = "error"
)

// InitLogger initializes the logger with level.
func InitLogger(cfg *config.Logger) (logger *zap.Logger, err error) {
	var level zapcore.Level

	switch cfg.Level {
	case DebugLevel, TraceLevel:
		level = zapcore.DebugLevel
	case InfoLevel:
		level = zapcore.InfoLevel
	case WarningLevel:
		level = zapcore.WarnLevel
	case ErrorLevel:
		level = zapcore.ErrorLevel
	default:
		return nil, errors.New("unknown logger level")
	}

	const EncodingJSON = "json"

	zapCfg := zap.NewDevelopmentConfig()
	zapCfg.DisableCaller = !cfg.EnableCaller
	zapCfg.DisableStacktrace = true
	zapCfg.Level = zap.NewAtomicLevelAt(level)

	if !cfg.HumanReadable {
		zapCfg.Encoding = EncodingJSON
	}

	logger, err = zapCfg.Build()
	if err != nil {
		return nil, fmt.Errorf("build: %w", err)
	}

	return logger.With(zap.String("level", cfg.Level)), nil
}
