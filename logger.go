package abing_logger

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var SugarLogger *zap.SugaredLogger

type Config struct {
	Filename   string
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Compress   bool
}

func Info(args ...interface{}) {
	if SugarLogger == nil {
		panic("please init log handler")
	}
	SugarLogger.Info(args)
}

func Infof(template string, args ...interface{}) {
	if SugarLogger == nil {
		panic("please init log handler")
	}
	SugarLogger.Infof(template, args)
}

func Warn(args ...interface{}) {
	if SugarLogger == nil {
		panic("please init log handler")
	}
	SugarLogger.Warn(args)
}

func Warnf(template string, args ...interface{}) {
	if SugarLogger == nil {
		panic("please init log handler")
	}
	SugarLogger.Warnf(template, args)
}

func Error(args ...interface{}) {
	if SugarLogger == nil {
		panic("please init log handler")
	}
	SugarLogger.Error(args)
}
func Errorf(template string, args ...interface{}) {
	if SugarLogger == nil {
		panic("please init log handler")
	}
	SugarLogger.Errorf(template, args)
}
func Panic(args ...interface{}) {
	if SugarLogger == nil {
		panic("please init log handler")
	}
	SugarLogger.Panic(args)
}
func Panicf(template string, args ...interface{}) {
	if SugarLogger == nil {
		panic("please init log handler")
	}
	SugarLogger.Panicf(template, args)
}

func InitLog(conf * Config){
	if conf == nil{
		panic("please init log handler")
	}
	writeSyncer := getLogWriter(conf)
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	logger := zap.New(core, zap.AddCaller())
	SugarLogger = logger.Sugar()
}
func CloseLog()error{
	return SugarLogger.Sync()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter(conf * Config) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   conf.Filename,
		MaxSize:    conf.MaxSize,
		MaxBackups: conf.MaxBackups,
		MaxAge:     conf.MaxAge,
		Compress:   conf.Compress,
	}
	return zapcore.AddSync(lumberJackLogger)
}
