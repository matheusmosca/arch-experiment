package xlog

import (
	"context"
	"io"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type xlogKey struct{}

func WithLogger(ctx context.Context, logger *zap.Logger) context.Context {
	return context.WithValue(ctx, xlogKey{}, logger)
}

func FromContext(ctx context.Context) *zap.Logger {
	logger, ok := ctx.Value(xlogKey{}).(*zap.Logger)
	if !ok {
		panic("no logger in context")
	}

	return logger
}

func New(w io.Writer) *zap.Logger {
	return zap.New(
		zapcore.NewCore(
			zapcore.NewJSONEncoder(zapcore.EncoderConfig{
				MessageKey:  "message",
				TimeKey:     "timestamp",
				LevelKey:    "level",
				EncodeLevel: zapcore.LowercaseColorLevelEncoder,
				EncodeTime:  zapcore.RFC3339NanoTimeEncoder,
			}),
			zapcore.AddSync(w),
			zapcore.InfoLevel,
		),
	)
}
