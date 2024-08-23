package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path/filepath"
	"pengyou/global/config"
)

var Logger *zap.Logger

// NewZapLogger creates a zap.Logger instance based on the Zap configuration.
func NewZapLogger(cfg *config.Zap) {
	// Set the Log level
	atomicLevel := zap.NewAtomicLevel()
	if err := atomicLevel.UnmarshalText([]byte(cfg.Level)); err != nil {
		return
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
			return
		}
		filePath := filepath.Join(cfg.Director, "app.Log")
		f, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			return
		}
		syncer = zapcore.AddSync(f)
	}

	// Create the core component
	core := zapcore.NewCore(encoder, syncer, atomicLevel)

	// Create the Logger
	Logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))

	Logger.Info("Logger initialized")
}
