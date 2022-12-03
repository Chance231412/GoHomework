package boot

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	g "homework_08/app/global"
	"os"
)

func LoggerSetup() {
	//自定义logger:
	//使用new(zapcore.core)来生成logger
	//core由三个部分组成:Encoder(日志格式)、LogLevel(日志等级)、WriteSyncer(日志写入地方)
	//配置LogLevel
	logLevel := zap.NewAtomicLevel()
	switch g.Config.Logger.LogLevel {
	case "info":
		logLevel.SetLevel(zap.InfoLevel)
	case "debug":
		logLevel.SetLevel(zap.DebugLevel)
	case "warn":
		logLevel.SetLevel(zap.WarnLevel)
	case "error":
		logLevel.SetLevel(zap.ErrorLevel)
	}

	//配置encoder,NewConsoleEncoder()返回的是一个更为平凡的encoder，就是一个毛坯，啥都要自己配,可以满足个性化需求
	//如果不指定对应key的name的话，对应key的信息不处理，即不会写入到文件中，如MessageKey为空的话，内容主体不处理，即看不到log内容
	encoder := zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		MessageKey:    "message",                        // 日志消息键
		LevelKey:      "level",                          // 日志等级键
		TimeKey:       "time",                           // 时间键
		NameKey:       "logger",                         // 日志记录器名
		CallerKey:     "caller",                         // 日志文件信息键
		StacktraceKey: "stacktrace",                     // 堆栈键
		LineEnding:    zapcore.DefaultLineEnding,        // 友好日志换行符
		EncodeLevel:   zapcore.CapitalColorLevelEncoder, // 友好日志等级名大小写（info INFO）
		//EncodeTime 值的类型是 TimeEncoder,也就是一个函数类型，用来将时间转换成人类可读(用户友好)的格式
		//type TimeEncoder func(time.Time, PrimitiveArrayEncoder)
		EncodeTime:     zapcore.ISO8601TimeEncoder,    // 友好日志时日期格式化
		EncodeDuration: zapcore.StringDurationEncoder, // 时间序列化
		EncodeCaller:   zapcore.FullCallerEncoder,     // 日志文件信息 short（包/文件.go:行号） full (文件位置.go:行号)
	})

	cores := [...]zapcore.Core{
		zapcore.NewCore(encoder, os.Stdin, logLevel),
		zapcore.NewCore(
			encoder,
			zapcore.AddSync(&lumberjack.Logger{
				Filename:   g.Config.Logger.SavePath,
				MaxSize:    g.Config.Logger.MaxSize,
				MaxAge:     g.Config.Logger.MaxAge,
				MaxBackups: g.Config.Logger.MaxBackups,
				LocalTime:  true,
				Compress:   g.Config.Logger.IsCompress,
			}),
			logLevel),
	}
	//使用core生成logger,core是一个接口类型，这里是将cores数组打散变为切片，然后转换成multicore类型，该类型实现了core接口
	g.Logger = zap.New(zapcore.NewTee(cores[:]...))

	defer func(logger *zap.Logger) {
		//同步调用底层核心的同步方法，刷新任何缓冲的日志条目。应用程序在退出之前应注意调用同步。
		_ = logger.Sync()
	}(g.Logger)

	g.Logger.Info("initialize logger successfully!")
}
