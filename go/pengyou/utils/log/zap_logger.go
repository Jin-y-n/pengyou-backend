package log

import (
	"os"
	"path/filepath"
	"pengyou/global/config"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

// NewZapLogger creates a zap.Logger instance based on the Zap configuration.
func NewZapLogger(cfg *config.Zap) (*zap.Logger, error) {
	// Set the log level
	atomicLevel := zap.NewAtomicLevel()
	if err := atomicLevel.UnmarshalText([]byte(cfg.Level)); err != nil {
		return nil, err
	}

	// Create the encoder
	encoder := cfg.Encoder()

	// Create the writer
	var syncer zapcore.WriteSyncer
	if cfg.LogInConsole {
		syncer = zapcore.AddSync(os.Stdout)
	} else {
		// Create the file path
		err := os.MkdirAll(cfg.Director, os.ModePerm)
		if err != nil {
			return nil, err
		}
		filePath := filepath.Join(cfg.Director, "app.log")
		f, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			return nil, err
		}
		syncer = zapcore.AddSync(f)
	}

	// Create the core component
	core := zapcore.NewCore(encoder, syncer, atomicLevel)

	// Create the logger
	Logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))

	return Logger, nil
}

// Info is a convenience function that wraps the Info method of the zap.Logger.
func Info(msg string, fields ...zap.Field) {
	Logger.Info(msg, fields...)
}

// Warn is a convenience function that wraps the Warn method of the zap.Logger.
func Warn(msg string, fields ...zap.Field) {
	Logger.Warn(msg, fields...)
}

// Error is a convenience function that wraps the Error method of the zap.Logger.
func Error(msg string, fields ...zap.Field) {
	Logger.Error(msg, fields...)
}

// Fatal is a convenience function that wraps the Fatal method of the zap.Logger.
func Fatal(msg string, fields ...zap.Field) {
	Logger.Fatal(msg, fields...)
}

// Debug is a convenience function that wraps the Debug method of the zap.Logger.
func Debug(msg string, fields ...zap.Field) {
	Logger.Debug(msg, fields...)
}
