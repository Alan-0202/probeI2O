package log

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// error logger
var zlogger *zap.SugaredLogger

var levelMap = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

func getLoggerLevel(lvl string) zapcore.Level {
	if level, ok := levelMap[lvl]; ok {
		return level
	}
	return zapcore.InfoLevel
}

func init() {
	fileName := "./log"

	level := getLoggerLevel("debug")
	syncWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:  fileName,
		MaxSize:   1 << 30, //1G
		LocalTime: true,
		Compress:  true,
	})
	encoder := zap.NewProductionEncoderConfig()
	encoder.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(zapcore.NewJSONEncoder(encoder), syncWriter, zap.NewAtomicLevelAt(level))
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	zlogger = logger.Sugar()
}

func Debug(args ...interface{}) {
	zlogger.Debug(args...)
}

func Debugf(template string, args ...interface{}) {
	zlogger.Debugf(template, args...)
}

func Info(args ...interface{}) {
	zlogger.Info(args...)
}

func Infof(template string, args ...interface{}) {
	zlogger.Infof(template, args...)
}

func Warn(args ...interface{}) {
	zlogger.Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	zlogger.Warnf(template, args...)
}

func Error(args ...interface{}) {
	zlogger.Error(args...)
}

func Errorf(template string, args ...interface{}) {
	zlogger.Errorf(template, args...)
}

func DPanic(args ...interface{}) {
	zlogger.DPanic(args...)
}

func DPanicf(template string, args ...interface{}) {
	zlogger.DPanicf(template, args...)
}

func Panic(args ...interface{}) {
	zlogger.Panic(args...)
}

func Panicf(template string, args ...interface{}) {
	zlogger.Panicf(template, args...)
}

func Fatal(args ...interface{}) {
	zlogger.Fatal(args...)
}

func Fatalf(template string, args ...interface{}) {
	zlogger.Fatalf(template, args...)
}

