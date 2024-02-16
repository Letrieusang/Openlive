package monitor

import (
	"api-openlive/config"
	"fmt"
	"github.com/getsentry/sentry-go"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
	"time"
)

// Logger wraps underlying logger
// The underlying logger can be either logrus or uber zap
type Logger struct {
	zapLogger *zap.SugaredLogger
}

// NewLogger creates new logger
func NewLogger(output ...string) Logger {
	zapConfig := zap.NewProductionConfig()
	tz := "Asia/Ho_Chi_Minh"

	loc, _ := time.LoadLocation(tz)
	zapConfig.EncoderConfig.EncodeTime = zapcore.TimeEncoder(func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.In(loc).Format("2006-01-02 15:04:05"))
	})
	zapConfig.DisableCaller = true
	err := sentry.Init(sentry.ClientOptions{
		Dsn: config.GetConfig().GetString("server.sentry"),
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	defer sentry.Flush(2 * time.Second)
	if len(output) != 0 {
		zapConfig.OutputPaths = output
	}
	z, _ := zapConfig.Build()
	return Logger{
		zapLogger: z.Sugar(),
	}
}

// Infof uses fmt.Sprintf to log a templated message.
func (l Logger) Infof(template string, args ...interface{}) {
	if len(args) == 0 {
		l.zapLogger.Info(template)
		return
	}
	l.zapLogger.Infof(template, args...)
}

// Panicf uses fmt.Sprintf to log a templated message.
func (l Logger) Panicf(template string, args ...interface{}) {
	if len(args) == 0 {
		l.zapLogger.Panic(template)
		return
	}
	l.zapLogger.Panicf(template, args...)
}

// Errorf uses fmt.Sprintf to log a templated message.
func (l Logger) Errorf(template string, args ...interface{}) {
	if len(args) == 0 {
		l.zapLogger.Error(template)
		return
	}
	l.zapLogger.Errorf(template, args...)
}

// Flush flushes any logs in buffer out
func (l Logger) Flush() {
	_ = l.zapLogger.Sync()
}

// WithTags set logging tags to logger
func (l Logger) WithTags(tags map[string]string) Logger {
	for k, v := range tags {
		l.zapLogger = l.zapLogger.With(k, v)
	}
	return l
}

func (l Logger) ErrofWithCapture(template string, args ...interface{}) {
	l.Errorf(template, args...)
	msg := fmt.Sprintf("Env: %s\nError: %s\n", os.Getenv("ENV"), template)
	sentry.CaptureMessage(fmt.Sprintf(msg, args...))
}

func (l Logger) InforfWithCapture(template string, args ...interface{}) {
	l.Infof(template, args...)
	msg := fmt.Sprintf("Env: %s\nMsg: %s\n", os.Getenv("ENV"), template)
	sentry.CaptureMessage(fmt.Sprintf(msg, args...))
}

func (l Logger) PanicfWithCapture(template string, panic ...interface{}) {
	l.Errorf(template)
	msg := fmt.Sprintf("Env: %s\nPanic: %s\n", os.Getenv("ENV"), template)
	sentry.CaptureMessage(fmt.Sprintf(msg, panic...))
}
