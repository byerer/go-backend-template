package logx

import (
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	rootLogger *Log
	logOnce    sync.Once
)

type Log struct {
	logger  *zap.Logger
	zapConf *zap.Config
	sinks   []Sink
}

type Option func(l *Log)

func SetUp(opts ...Option) *Log {
	logOnce.Do(func() {
		rootLogger = &Log{}
		for _, opt := range opts {
			opt(rootLogger)
		}
		if rootLogger.zapConf == nil {
			rootLogger.zapConf = defaultConfig()
		}
		log, err := rootLogger.zapConf.Build(zap.AddCallerSkip(1), zap.AddStacktrace(zapcore.DPanicLevel))
		if err != nil {
			panic(err)
		}
		rootLogger.logger = log
	})
	return rootLogger
}

func defaultConfig() *zap.Config {
	return &zap.Config{
		Level:            zap.NewAtomicLevelAt(zapcore.InfoLevel),
		Development:      false,
		Sampling:         nil,
		Encoding:         "console",
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "ts",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
	}
}

func SetLevel(level zapcore.Level) {
	rootLogger.zapConf.Level.SetLevel(level)
}

func Debug(v ...any) {
	rootLogger.logger.Sugar().Debug(v...)
}

func Info(v ...any) {
	rootLogger.logger.Sugar().Info(v...)
}

func Warn(v ...any) {
	rootLogger.logger.Sugar().Warn(v...)
}

func Error(v ...any) {
	rootLogger.logger.Sugar().Error(v...)
}

func Panic(v ...any) {
	rootLogger.logger.Sugar().Panic(v...)
}

func Fatal(v ...any) {
	rootLogger.logger.Sugar().Fatal(v...)
}

func Debugf(format string, v ...any) {
	rootLogger.logger.Sugar().Debugf(format, v...)
}

func Infof(format string, v ...any) {
	rootLogger.logger.Sugar().Infof(format, v...)
}

func Warnf(format string, v ...any) {
	rootLogger.logger.Sugar().Warnf(format, v...)
}

func Errorf(format string, v ...any) {
	rootLogger.logger.Sugar().Errorf(format, v...)
}

func Panicf(format string, v ...any) {
	rootLogger.logger.Sugar().Panicf(format, v...)
}

func Fatalf(format string, v ...any) {
	rootLogger.logger.Sugar().Fatalf(format, v...)
}

func Debugw(msg string, keysAndValues ...any) {
	rootLogger.logger.Sugar().Debugw(msg, keysAndValues...)
}

func Infow(msg string, keysAndValues ...any) {
	rootLogger.logger.Sugar().Infow(msg, keysAndValues...)
}

func Warnw(msg string, keysAndValues ...any) {
	rootLogger.logger.Sugar().Warnw(msg, keysAndValues...)
}

func Errorw(msg string, keysAndValues ...any) {
	rootLogger.logger.Sugar().Errorw(msg, keysAndValues...)
}

func Panicw(msg string, keysAndValues ...any) {
	rootLogger.logger.Sugar().Panicw(msg, keysAndValues...)
}

func Fatalw(msg string, keysAndValues ...any) {
	rootLogger.logger.Sugar().Fatalw(msg, keysAndValues...)
}
