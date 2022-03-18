package wx

import (
	"context"
	"os"
	"time"

	"github.com/natefinch/lumberjack"
	"github.com/tidwall/pretty"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// LoggerConfig keeps the settings to configure logger.
type LoggerConfig struct {
	// Filename is the file to write logs to.
	Filename string `json:"filename"`

	// Options optional settings to configure logger.
	Options *LoggerOptions `json:"options"`
}

// LoggerOptions optional settings to configure logger.
type LoggerOptions struct {
	// MaxSize is the maximum size in megabytes of the log file before it gets
	// rotated. It defaults to 100 megabytes.
	MaxSize int `json:"max_size"`

	// MaxAge is the maximum number of days to retain old log files based on the
	// timestamp encoded in their filename.  Note that a day is defined as 24
	// hours and may not exactly correspond to calendar days due to daylight
	// savings, leap seconds, etc. The default is not to remove old log files
	// based on age.
	MaxAge int `json:"max_age"`

	// MaxBackups is the maximum number of old log files to retain. The default
	// is to retain all old log files (though MaxAge may still cause them to get
	// deleted.)
	MaxBackups int `json:"max_backups"`

	// Compress determines if the rotated log files should be compressed
	// using gzip. The default is not to perform compression.
	Compress bool `json:"compress"`

	// Stderr specifies the stderr for logger
	Stderr bool `json:"stderr"`

	// ZapOptions specifies the zap options stderr for logger
	ZapOptions []zap.Option `json:"zap_options"`
}

// newLogger returns a new logger.
func newLogger(cfg *LoggerConfig) *zap.Logger {
	if len(cfg.Filename) == 0 {
		return debugLogger(cfg.Options.ZapOptions...)
	}

	c := zap.NewProductionEncoderConfig()

	c.TimeKey = "time"
	c.EncodeTime = MyTimeEncoder
	c.EncodeCaller = zapcore.FullCallerEncoder

	ws := make([]zapcore.WriteSyncer, 0, 2)

	ws = append(ws, zapcore.AddSync(&lumberjack.Logger{
		Filename:   cfg.Filename,
		MaxSize:    cfg.Options.MaxSize,
		MaxAge:     cfg.Options.MaxAge,
		MaxBackups: cfg.Options.MaxBackups,
		Compress:   cfg.Options.Compress,
		LocalTime:  true,
	}))

	if cfg.Options.Stderr {
		ws = append(ws, zapcore.Lock(os.Stderr))
	}

	core := zapcore.NewCore(zapcore.NewJSONEncoder(c), zapcore.NewMultiWriteSyncer(ws...), zap.DebugLevel)

	return zap.New(core, cfg.Options.ZapOptions...)
}

func debugLogger(options ...zap.Option) *zap.Logger {
	cfg := zap.NewDevelopmentConfig()

	cfg.DisableCaller = true
	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	cfg.EncoderConfig.EncodeTime = MyTimeEncoder
	cfg.EncoderConfig.EncodeCaller = zapcore.FullCallerEncoder

	l, _ := cfg.Build(options...)

	return l
}

// MyTimeEncoder zap time encoder.
func MyTimeEncoder(t time.Time, e zapcore.PrimitiveArrayEncoder) {
	e.AppendString(t.In(time.FixedZone("CST", 8*3600)).Format("2006-01-02 15:04:05"))
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

type LogFn func(ctx context.Context, logger *zap.Logger, data *LogData)

type wxlogger struct {
	zlog  *zap.Logger
	logFn LogFn
}

func (l *wxlogger) Log(ctx context.Context, data *LogData) {
	if l.logFn != nil {
		l.logFn(ctx, l.zlog, data)

		return
	}

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

// NewLogger returns new logger
func NewLogger(cfg *LoggerConfig, fn LogFn) Logger {
	return &wxlogger{
		zlog:  newLogger(cfg),
		logFn: fn,
	}
}

// DefaultLogger returns default logger
func DefaultLogger() Logger {
	return &wxlogger{
		zlog: debugLogger(),
	}
}
