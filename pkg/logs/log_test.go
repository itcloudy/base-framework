// Copyright 2018 cloudy 272685110@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package logs

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"io/ioutil"
	"os"
	"path"
	"testing"
)

func TestInitLogger(t *testing.T) {
	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})
	// 打印所有级别的日志
	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.DebugLevel
	})
	cwd, _ := os.Getwd()
	hook := lumberjack.Logger{
		Filename:   path.Join(cwd, "abc.log"),
		MaxSize:    1024, // megabytes
		MaxBackups: 3,
		MaxAge:     7,    //days
		Compress:   true, // disabled by default
	}

	topicErrors := zapcore.AddSync(ioutil.Discard)
	fileWriter := zapcore.AddSync(&hook)

	// High-priority output should also go to standard error, and low-priority
	// output should also go to standard out.
	consoleDebugging := zapcore.Lock(os.Stdout)

	// Optimize the Kafka output for machine consumption and the console output
	// for human operators.
	kafkaEncoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())

	// Join the outputs, encoders, and level-handling functions into
	// zapcore.Cores, then tee the four cores together.
	kfk := zapcore.NewCore(kafkaEncoder, topicErrors, highPriority)
	consl := zapcore.NewCore(consoleEncoder, consoleDebugging, lowPriority)
	fileC := zapcore.NewCore(consoleEncoder, fileWriter, highPriority)

	core := zapcore.NewTee(
		// 打印在kafka topic中（伪造的case）
		kfk,
		// 打印在控制台
		consl,
		// 打印在文件中
		fileC,
	)

	// From a zapcore.Core, it's easy to construct a Logger.
	logger := zap.New(core)
	defer logger.Sync()
	logger.Error("constructed a info logger", zap.Int("test", 1))
	logger.Error("constructed a error logger", zap.Int("test", 2))

}
func tst() {
	Logger.Error("error", zap.String("test_key", "kafka test value"))
	Logger.Error("error", zap.String("test_key", "kafka test value"))
	Logger.Error("error", zap.String("test_key", "kafka test value"))
	Logger.Error("error", zap.String("test_key", "kafka test value"))
	Logger.Error("error", zap.String("test_key", "kafka test value"))
	Logger.Error("error", zap.String("test_key", "kafka test value"))

}
