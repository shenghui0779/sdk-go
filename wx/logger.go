package wx

import (
	"context"
	"strings"
	"time"

	"github.com/shenghui0779/yiigo"
	"go.uber.org/zap"
)

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

var replacer = strings.NewReplacer("\n", "", "\t", "", "\r", "#")

type wxlogger struct{}

func (l *wxlogger) Log(ctx context.Context, data *LogData) {
	fields := make([]zap.Field, 0, 7)

	fields = append(fields,
		zap.String("method", data.Method),
		zap.String("url", data.URL),
		zap.String("body", replacer.Replace(string(data.Body))),
		zap.String("response", replacer.Replace(string(data.Response))),
		zap.Int("status", data.StatusCode),
		zap.String("duration", data.Duration.String()),
	)

	if data.Error != nil {
		fields = append(fields, zap.Error(data.Error))

		yiigo.Logger().Error("[gochat] action do error", fields...)

		return
	}

	yiigo.Logger().Info("[gochat] action do info", fields...)
}

// DefaultLogger returns default logger
func DefaultLogger() Logger {
	return new(wxlogger)
}
