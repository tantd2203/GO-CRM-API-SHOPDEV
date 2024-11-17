package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func main() {

	//sugar := zap.NewExample().Sugar()
	//sugar.Info("Hello name :%s", "tipsgo")
	////
	//
	//logger := zap.NewExample()
	//logger.Info("Hello", zap.String("name", "tips js"))

	//logger := zap.NewExample()
	//logger.Info("Hello")
	//
	//// Development
	//logger, _ = zap.NewDevelopment()
	//logger.Info("NewDevelopment")
	//
	////production
	//logger, _ = zap.NewProduction()
	//logger.Info("NewProduction")

	// 3.

	encoder := getEncoderLog()
	sync := getWriterSync()
	core := zapcore.NewCore(encoder, sync, zapcore.InfoLevel)
	logger := zap.New(core, zap.AddCaller())

	logger.Info("Info log", zap.Int("line", 1))
	logger.Error("Error log", zap.Int("line", 2))

}

// format log
func getEncoderLog() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig()
	//	ts":1731859099.214258 ->2024-11-17T22:58:19.213+0700
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	//ts -Time
	encodeConfig.TimeKey = "time"
	// from info
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	// "caller":"cli/main.log.go:22
	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder

	return zapcore.NewJSONEncoder(encodeConfig)

}

//

func getWriterSync() zapcore.WriteSyncer {
	// permission ModePerm
	file, _ := os.OpenFile("./log/log.txt", os.O_CREATE|os.O_WRONLY, os.ModePerm)

	syncFile := zapcore.AddSync(file)

	syncConsole := zapcore.AddSync(os.Stderr)

	return zapcore.NewMultiWriteSyncer(syncConsole, syncFile)
}
