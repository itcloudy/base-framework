// Copyright 2018 cloudy 272685110@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package logs

import (
	"fmt"
	"github.com/Shopify/sarama"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

var Logger *zap.Logger

type LogKafka struct {
	Producer sarama.SyncProducer
	Topic    string
}

func (lk *LogKafka) Write(p []byte) (n int, err error) {
	msg := &sarama.ProducerMessage{}
	msg.Topic = lk.Topic
	msg.Value = sarama.ByteEncoder(p)
	_, _, err = lk.Producer.SendMessage(msg)
	if err != nil {
		return
	}
	return

}
func InitLogger(mode string, fileName string, maxSize, maxBackups, maxAge int, compress bool, enableKafka bool, kafkaAddress []string) {
	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})
	// 打印所有级别的日志
	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.DebugLevel
	})
	var allCore []zapcore.Core

	hook := lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    maxSize, // megabytes
		MaxBackups: maxBackups,
		MaxAge:     maxAge,   //days
		Compress:   compress, // disabled by default
	}

	fileWriter := zapcore.AddSync(&hook)

	// High-priority output should also go to standard error, and low-priority
	// output should also go to standard out.
	consoleDebugging := zapcore.Lock(os.Stdout)

	// for human operators.
	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())

	// Join the outputs, encoders, and level-handling functions into
	// zapcore.Cores, then tee the four cores together.
	// kafka
	if len(kafkaAddress) > 0 && enableKafka {
		var (
			kl  LogKafka
			err error
		)
		kl.Topic = "go_framework_log"
		// 设置日志输入到Kafka的配置
		config := sarama.NewConfig()
		//等待服务器所有副本都保存成功后的响应
		config.Producer.RequiredAcks = sarama.WaitForAll
		//随机的分区类型
		config.Producer.Partitioner = sarama.NewRandomPartitioner
		//是否等待成功和失败后的响应,只有上面的RequireAcks设置不是NoReponse这里才有用.
		config.Producer.Return.Successes = true
		config.Producer.Return.Errors = true

		kl.Producer, err = sarama.NewSyncProducer(kafkaAddress, config)
		if err != nil {
			fmt.Printf("connect kafka failed: %+v\n", err)
			os.Exit(-1)
		}
		topicErrors := zapcore.AddSync(&kl)
		// 打印在kafka
		kafkaEncoder := zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig())
		var kafkaCore zapcore.Core
		if mode == "debug" {
			kafkaCore = zapcore.NewCore(kafkaEncoder, topicErrors, lowPriority)

		} else {
			kafkaCore = zapcore.NewCore(kafkaEncoder, topicErrors, highPriority)

		}
		allCore = append(allCore, kafkaCore)
	}
	if mode == "debug" {
		allCore = append(allCore, zapcore.NewCore(consoleEncoder, consoleDebugging, lowPriority))
		allCore = append(allCore, zapcore.NewCore(consoleEncoder, fileWriter, lowPriority))
	} else {
		allCore = append(allCore, zapcore.NewCore(consoleEncoder, consoleDebugging, highPriority))
		allCore = append(allCore, zapcore.NewCore(consoleEncoder, fileWriter, highPriority))
	}
	core := zapcore.NewTee(allCore...)

	// From a zapcore.Core, it's easy to construct a Logger.
	Logger = zap.New(core)
}
