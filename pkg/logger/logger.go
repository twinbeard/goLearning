package logger

import (
	"os"

	"github.com/natefinch/lumberjack"
	"github.com/twinbeard/goLearning/setting"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LoggerZap struct {
	*zap.Logger
}

func NewLogger(config setting.LoggerSetting) *LoggerZap {
	logLevel := config.Log_level
	// debug-> info-> warn ->error->fatal->panic
	var level zapcore.Level
	switch logLevel {
	case "debug":
		level = zapcore.DebugLevel
	case "info":
		level = zapcore.InfoLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	default:
		level = zapcore.InfoLevel
	}
	encoder := getEncoderLog() // custom log
	hook := lumberjack.Logger{
		Filename:   config.File_log_name, // tên log file
		MaxSize:    config.Max_size,      //  dung lượng file log - 500MB
		MaxBackups: config.Max_backups,   // số lượng file backup - 3
		MaxAge:     config.Max_Age,       //days - // số ngày file log được giữ lại
		Compress:   config.Compress,      // disabled by default - nén file log
	} // write to file and console
	core := zapcore.NewCore(encoder, //
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)), // write to file
		level) // log level
	logger := zap.New(core, zap.AddCaller()) // add caller (thêm thông tin file gọi đến log )
	// logger.Info("Info log", zap.Int("line", 1))
	// logger.Error("Error log", zap.Int("line", 2))
	return &LoggerZap{logger} // add caller
}

// format log message
func getEncoderLog() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig() // default config

	// 1716714967.877995 -> 2024-05-26T16:16:07.877+0700
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// ts -> time
	encodeConfig.TimeKey = "time"
	// from info INFO
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	//"caller":"cli/main.log.go:24"
	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encodeConfig)
}

func getWriterSync() zapcore.WriteSyncer {
	file, _ := os.OpenFile("./log/log.txt", os.O_CREATE|os.O_WRONLY, os.ModePerm)
	syncFile := zapcore.AddSync(file)
	syncConsole := zapcore.AddSync(os.Stderr)
	return zapcore.NewMultiWriteSyncer(syncConsole, syncFile)
}
