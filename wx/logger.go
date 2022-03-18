package wx

import (
	"context"
	"time"

	"github.com/tidwall/pretty"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func debugLogger(options ...zap.Option) *zap.Logger {
	cfg := zap.NewDevelopmentConfig()

	cfg.DisableCaller = true
	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	cfg.EncoderConfig.EncodeTime = func(t time.Time, pae zapcore.PrimitiveArrayEncoder) {
		pae.AppendString(t.In(time.FixedZone("CST", 8*3600)).Format("2006-01-02 15:04:05"))
	}
	cfg.EncoderConfig.EncodeCaller = zapcore.FullCallerEncoder

	l, _ := cfg.Build(options...)

	return l
}

type Logger interface {
	Log(ctx context.Context, data *LogData)
}

type LogData struct {
	URL        string        `json:"url"`
	Method     string        `json:"method"`
	Body       []byte        `json:"body"`
	StatusCode int           `json:"status_code"`
	Response   []byte        `json:"response"`
	Duration   time.Duration `json:"duration"`
	Error      error         `json:"error"`
}

type wxlogger struct {
	zlog *zap.Logger
}

func (l *wxlogger) Log(ctx context.Context, data *LogData) {
	fields := make([]zap.Field, 0, 7)

	fields = append(fields,
		zap.String("method", data.Method),
		zap.String("url", data.URL),
		zap.ByteString("body", pretty.Ugly(data.Body)),
		zap.ByteString("response", pretty.Ugly(data.Response)),
		zap.Int("status", data.StatusCode),
		zap.String("duration", data.Duration.String()),
	)

	if data.Error != nil {
		fields = append(fields, zap.Error(data.Error))

		l.zlog.Error("[gochat] action error", fields...)

		return
	}

	l.zlog.Info("[gochat] action info", fields...)
}

// DefaultLogger returns default logger
func DefaultLogger() Logger {
	return &wxlogger{
		zlog: debugLogger(),
	}
}
