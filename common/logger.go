package common

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger это алиас для zap.Logger.
// Введен, чтобы не тянуть зависимости в другие места
type Logger = zap.Logger

// В отдельное место вытащил настройки для логгера
var (
	encoderConfig = zapcore.EncoderConfig{
		TimeKey:     "time",
		LevelKey:    "level",
		MessageKey:  "message",
		EncodeLevel: levelEncoder,
		EncodeTime:  timeEncoder,
	}
)

// NewLogger возвращает экземпляр логгера
func NewLogger(config *Config) (*Logger, error) {
	level := zap.NewAtomicLevel()

	if config.Server.Debug {
		level.SetLevel(zapcore.DebugLevel)
	}

	logConfig := zap.Config{
		Encoding:      "console",
		Level:         level,
		OutputPaths:   []string{"stdout"},
		EncoderConfig: encoderConfig,
	}

	log, err := logConfig.Build()
	if err != nil {
		return nil, err
	}

	log.Debug("Enabled DEBUG mode")

	return log, nil
}

func levelEncoder(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + l.String() + "]")
}

func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("02/01/2006 15:04:05.99"))
}
