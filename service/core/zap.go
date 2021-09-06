package core

import (
	"Hyaenidae/global"
	"Hyaenidae/utils"
	"fmt"
	"os"
	"path"
	"time"

	zaprotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var level zapcore.Level

//@function: Zap
//@description: 初始化Zap相关配置
//@return: *zap.Logger

func Zap() (logger *zap.Logger) {
	if ok, _ := utils.IsExists(global.Hyaenidae_CONFIG.Zap.Director); !ok { //判断配置文件设置的日志文件夹是否存在，不存在则创建
		fmt.Printf("create %v directory\n", global.Hyaenidae_CONFIG.Zap.Director)
		os.Mkdir(global.Hyaenidae_CONFIG.Zap.Director, os.ModePerm)
	}

	switch global.Hyaenidae_CONFIG.Zap.Level { // 初始化配置文件的Level,默认为Info
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	case "dpanic":
		level = zap.DPanicLevel
	case "panic":
		level = zap.PanicLevel
	case "fatal":
		level = zap.FatalLevel
	default:
		level = zap.InfoLevel
	}

	if level == zap.DebugLevel || level == zap.ErrorLevel {
		logger = zap.New(getEncoderCore(), zap.AddStacktrace(level))
	} else {
		logger = zap.New(getEncoderCore())
	}
	if global.Hyaenidae_CONFIG.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}

//@function: getEncoderConfig
//@description: 获取zapcore.EncoderConfig
//@return: zapcore.EncoderConfig

func getEncoderConfig() (config zapcore.EncoderConfig) {
	config = zapcore.EncoderConfig{
		MessageKey:     "message", //日志内容对应的key名，此参数必须不为空，否则日志主体不处理
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  global.Hyaenidae_CONFIG.Zap.StacktraceKey,
		LineEnding:     zapcore.DefaultLineEnding,      //默认换行
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  //小写
		EncodeTime:     CustomTimeEncoder,              //自定义时间显示
		EncodeDuration: zapcore.SecondsDurationEncoder, //日期转为s
		EncodeCaller:   zapcore.FullCallerEncoder,      //记录调用路径格式
	}
	switch {
	case global.Hyaenidae_CONFIG.Zap.EncodeLevel == "LowercaseLevelEncoder": // 小写编码器(默认)
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	case global.Hyaenidae_CONFIG.Zap.EncodeLevel == "LowercaseColorLevelEncoder": // 小写编码器带颜色
		config.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	case global.Hyaenidae_CONFIG.Zap.EncodeLevel == "CapitalLevelEncoder": // 大写编码器
		config.EncodeLevel = zapcore.CapitalLevelEncoder
	case global.Hyaenidae_CONFIG.Zap.EncodeLevel == "CapitalColorLevelEncoder": // 大写编码器带颜色
		config.EncodeLevel = zapcore.CapitalColorLevelEncoder
	default:
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	}
	return config
}

//@function: getEncoder
//@description: 获取zapcore.Encoder
//@return: zapcore.Encoder

func getEncoder() zapcore.Encoder {
	if global.Hyaenidae_CONFIG.Zap.Format == "json" {
		return zapcore.NewJSONEncoder(getEncoderConfig())
	}
	return zapcore.NewConsoleEncoder(getEncoderConfig())
}

//@function: getEncoderCore
//@description: 获取zapcore.Core
//@return: zapcore.Core

func getEncoderCore() (core zapcore.Core) {
	writer, err := GetWriteSyncer() // 使用file-rotatelogs进行日志分割
	if err != nil {
		fmt.Printf("Get Write Syncer Failed err:%v", err.Error())
		return
	}
	return zapcore.NewCore(getEncoder(), writer, level)
}

//@function: CustomTimeEncoder
//@description: 设置日志自定义事件
//@return: enc.AppendString

func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(global.Hyaenidae_CONFIG.Zap.Prefix + "2006/01/02 - 15:04:05.000"))
}

//@function: GetWriteSyncer
//@description: 利用file-rotatelogs实现日志分割
//@return: zapcore.AddSync(fileWriter)

func GetWriteSyncer() (zapcore.WriteSyncer, error) {
	fileWriter, err := zaprotatelogs.New(
		path.Join(global.Hyaenidae_CONFIG.Zap.Director, "%Y-%m-%d.log"),
		zaprotatelogs.WithMaxAge(7*24*time.Hour),     //日志保存7天
		zaprotatelogs.WithRotationTime(24*time.Hour), //以天为单位进行分割
	)
	if global.Hyaenidae_CONFIG.Zap.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter)), err
	}
	return zapcore.AddSync(fileWriter), err
}
