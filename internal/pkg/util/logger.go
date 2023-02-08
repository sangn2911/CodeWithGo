package util

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/labstack/echo"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.SugaredLogger

func InitializeLogger() {
	if Logger != nil {
		return
	}
	cfg := zap.Config{
		Encoding:    "json",
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		OutputPaths: []string{"stderr", "log/log.log"},

		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:  "message",
			TimeKey:     "time",
			LevelKey:    "level",
			EncodeLevel: CustomLevelEncoder,
			EncodeTime:  SyslogTimeEncoder,
			EncodeName:  NameEncoder,
		},
	}
	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	Logger = logger.Sugar()
}

func SyslogTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}

func CustomLevelEncoder(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + level.String() + "]")
}

func NameEncoder(name string, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + name + "]")
}

func MuxLoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		Logger.Infof("method=%v route=%v", r.Method, r.URL.Path)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func EchoLoggingMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return echo.HandlerFunc(func(ctx echo.Context) error {
		Logger.Infof("method=%v route=%v", ctx.Request().Method, ctx.Request().URL.Path)
		return next(ctx)
	})
}

func FiberLoggingMiddleware(ctx *fiber.Ctx) error {
	Logger.Infof("method=%v route=%v", ctx.Method(), ctx.Path())
	return ctx.Next()
}
