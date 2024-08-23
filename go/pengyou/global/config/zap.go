package config

import (
	"go.uber.org/zap/zapcore"
	"time"
)

// Zap holds configuration settings for the zap logger.
type Zap struct {
	Level         string `mapstructure:"level" json:"level" yaml:"level"`                            // Level
	Prefix        string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`                         // Prefix
	Format        string `mapstructure:"format" json:"format" yaml:"format"`                         // Output format
	Director      string `mapstructure:"director" json:"director" yaml:"director"`                   // Directory
	EncodeLevel   string `mapstructure:"encode-level" json:"encode-level" yaml:"encode-level"`       // Level encoding
	StacktraceKey string `mapstructure:"stacktrace-key" json:"stacktrace-key" yaml:"stacktrace-key"` // Stacktrace key
	ShowLine      bool   `mapstructure:"show-line" json:"show-line" yaml:"show-line"`                // Show line number
	LogInConsole  bool   `mapstructure:"log-in-console" json:"log-in-console" yaml:"log-in-console"` // log to console
	RetentionDay  int    `mapstructure:"retention-day" json:"retention-day" yaml:"retention-day"`    // log retention days
}

// Levels converts the string level to a slice of zapcore.Level.
func (c *Zap) Levels() []zapcore.Level {
	levels := make([]zapcore.Level, 0, 7)
	level, err := zapcore.ParseLevel(c.Level)
	if err != nil {
		level = zapcore.DebugLevel
	}
	for ; level <= zapcore.FatalLevel; level++ {
		levels = append(levels, level)
	}
	return levels
}

// Encoder creates a zap core.Encoder based on the configuration.
func (c *Zap) Encoder() zapcore.Encoder {
	config := zapcore.EncoderConfig{
		TimeKey:       "time",
		NameKey:       "name",
		LevelKey:      "level",
		CallerKey:     "caller",
		MessageKey:    "message",
		StacktraceKey: c.StacktraceKey,
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeTime: func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
			encoder.AppendString(c.Prefix + t.Format("2006-01-02 15:04:05.000"))
		},
		EncodeLevel:    c.LevelEncoder(),
		EncodeCaller:   zapcore.FullCallerEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
	}
	if c.Format == "json" {
		return zapcore.NewJSONEncoder(config)
	}
	return zapcore.NewConsoleEncoder(config)
}

// LevelEncoder returns a zapcore.LevelEncoder based on the EncodeLevel setting.
func (c *Zap) LevelEncoder() zapcore.LevelEncoder {
	switch c.EncodeLevel {
	case "LowercaseLevelEncoder": // Lowercase encoder (default)
		return zapcore.LowercaseLevelEncoder
	case "LowercaseColorLevelEncoder": // Lowercase encoder with color
		return zapcore.LowercaseColorLevelEncoder
	case "CapitalLevelEncoder": // Uppercase encoder
		return zapcore.CapitalLevelEncoder
	case "CapitalColorLevelEncoder": // Uppercase encoder with color
		return zapcore.CapitalColorLevelEncoder
	default:
		return zapcore.LowercaseLevelEncoder
	}
}
