// Copyright 2018 cloudy itcloudy@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package cmd

import (
	"fmt"
	"github.com/itcloudy/base-framework/pkg/conf"
	"github.com/itcloudy/base-framework/pkg/consts"
	"github.com/itcloudy/base-framework/pkg/logs"
	"github.com/olivere/elastic"
	"go.uber.org/zap"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"path/filepath"
	"strings"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Initial config generation",
	Run: func(cmd *cobra.Command, args []string) {
		// 指定配置文件位置
		configPath, _ := cmd.Flags().GetString("path")
		err := conf.FillRuntimePaths()
		if err != nil {
			fmt.Printf("Filling config: %+v\n", err)

			logs.Logger.Fatal("Filling config", zap.Error(err))
		}

		if configPath == "" {
			configPath = filepath.Join("config", consts.DefaultConfigFile)

		}
		err = viper.Unmarshal(&conf.Config)
		if err != nil {
			fmt.Printf("Marshalling config to global struct variable: %+v\n", err)

		}

		err = conf.SaveConfig(configPath)
		if err != nil {
			fmt.Printf("Saving config failed: %+v\n", err)
		}
		fmt.Println("config file is saved success and path is " + configPath)

	},
}

func init() {
	viper.SetEnvPrefix("GO_FRAMEWORK")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// run mode
	configCmd.Flags().String("mode", "debug", "server run mode")
	viper.BindPFlag("Mode", configCmd.Flags().Lookup("mode"))
	//token
	configCmd.Flags().DurationVar(&conf.Config.TokenExpire, "token-expire", 2*time.Hour, "Token Expire ")
	viper.BindPFlag("TokenExpire", configCmd.Flags().Lookup("token-expire"))
	// Command flags
	configCmd.Flags().String("path", "", "Generate config to (default config/config.toml)")

	// super user
	configCmd.Flags().StringVar(&conf.Config.Admin.UserName, "admin-name", "admin", "super user name")
	configCmd.Flags().StringVar(&conf.Config.Admin.Password, "admin-pwd", "itcloudy", "super user password")
	configCmd.Flags().StringVar(&conf.Config.Admin.Email, "admin-email", "itcloudy@qq.com", "super user email")
	configCmd.Flags().StringVar(&conf.Config.Admin.Mobile, "admin-mobile", "13888888888", "super user mobile")
	viper.BindPFlag("Admin.UserName", configCmd.Flags().Lookup("admin-name"))
	viper.BindPFlag("Admin.Password", configCmd.Flags().Lookup("admin-pwd"))
	viper.BindPFlag("Admin.Email", configCmd.Flags().Lookup("admin-email"))
	viper.BindPFlag("Admin.Mobile", configCmd.Flags().Lookup("admin-mobile"))

	// jwt
	configCmd.Flags().StringVar(&conf.Config.JwtPrivatePath, "jwt-pri", "", "jwt private path")
	configCmd.Flags().StringVar(&conf.Config.JwtPublicPath, "jwt-pub", "", "jwt pblic path")
	viper.BindPFlag("JwtPrivatePath", configCmd.Flags().Lookup("jwt-pri"))
	viper.BindPFlag("JwtPublicPath", configCmd.Flags().Lookup("jwt-pub"))

	// HTTP Server
	configCmd.Flags().StringVar(&conf.Config.HTTP.Host, "httpHost", "127.0.0.1", " HTTP server host")
	configCmd.Flags().IntVar(&conf.Config.HTTP.Port, "httpPort", 8000, " HTTP server port")
	configCmd.Flags().BoolVar(&conf.Config.RPC.Enable, "enableHttp", true, "rpc server enable")
	viper.BindPFlag("HTTP.Host", configCmd.Flags().Lookup("httpHost"))
	viper.BindPFlag("HTTP.Port", configCmd.Flags().Lookup("httpPort"))
	viper.BindPFlag("HTTP.Enable", configCmd.Flags().Lookup("enableHttp"))

	// RPC Server
	configCmd.Flags().StringVar(&conf.Config.RPC.Host, "rpcHost", "127.0.0.1", "RPC server host")
	configCmd.Flags().IntVar(&conf.Config.RPC.Port, "rpcPort", 9000, "RPC server port")
	configCmd.Flags().BoolVar(&conf.Config.RPC.Enable, "rpcEnable", true, "rpc server enable")
	viper.BindPFlag("RPC.Host", configCmd.Flags().Lookup("rpcHost"))
	viper.BindPFlag("RPC.Port", configCmd.Flags().Lookup("rpcPort"))
	viper.BindPFlag("RPC.Enable", configCmd.Flags().Lookup("rpcEnable"))

	// DB
	configCmd.Flags().StringVar(&conf.Config.DB.Host, "dbHost", "127.0.0.1", "DB host")
	configCmd.Flags().IntVar(&conf.Config.DB.Port, "dbPort", 5432, "DB port")
	configCmd.Flags().StringVar(&conf.Config.DB.Name, "dbName", "go", "DB name")
	configCmd.Flags().StringVar(&conf.Config.DB.User, "dbUser", "postgres", "DB username")
	configCmd.Flags().StringVar(&conf.Config.DB.Password, "dbPassword", "postgres", "DB password")
	configCmd.Flags().IntVar(&conf.Config.DB.LockTimeout, "dbLockTimeout", 5000, "DB lock timeout")
	configCmd.Flags().StringVar(&conf.Config.DB.DbType, "dbType", consts.DefaultDatabase, "db type :mysql/postgres default postgres")
	configCmd.Flags().StringVar(&conf.Config.DB.Charset, "charset", "utf8", "db charset only for mysql")

	viper.BindPFlag("DB.Name", configCmd.Flags().Lookup("dbName"))
	viper.BindPFlag("DB.Host", configCmd.Flags().Lookup("dbHost"))
	viper.BindPFlag("DB.Port", configCmd.Flags().Lookup("dbPort"))
	viper.BindPFlag("DB.User", configCmd.Flags().Lookup("dbUser"))
	viper.BindPFlag("DB.Password", configCmd.Flags().Lookup("dbPassword"))
	viper.BindPFlag("DB.Charset", configCmd.Flags().Lookup("charset"))
	viper.BindPFlag("DB.DbType", configCmd.Flags().Lookup("dbType"))
	viper.BindPFlag("DB.LockTimeout", configCmd.Flags().Lookup("dbLockTimeout"))

	// Centrifugo
	configCmd.Flags().BoolVar(&conf.Config.Centrifugo.Enable, "centEnable", false, "Centrifugo enable")
	configCmd.Flags().StringVar(&conf.Config.Centrifugo.Secret, "centSecret", "127.0.0.1", "Centrifugo secret")
	configCmd.Flags().StringVar(&conf.Config.Centrifugo.URL, "centUrl", "127.0.0.1", "Centrifugo URL")
	viper.BindPFlag("Centrifugo.Enable", configCmd.Flags().Lookup("centEnable"))
	viper.BindPFlag("Centrifugo.Secret", configCmd.Flags().Lookup("centSecret"))
	viper.BindPFlag("Centrifugo.URL", configCmd.Flags().Lookup("centUrl"))

	// Log
	configCmd.Flags().StringSliceVar(&conf.Config.Log.KafkaAddress, "kafkaAddress", []string{"localhost:9092"}, "kafka address for receive log")
	configCmd.Flags().StringVar(&conf.Config.Log.FileName, "logFile", "", "log file name")
	configCmd.Flags().IntVar(&conf.Config.Log.MaxSize, "logFileMaxSize", 1024, "max log file size")
	configCmd.Flags().IntVar(&conf.Config.Log.MaxBackups, "logFileBackups", 3, "number of log backup")
	configCmd.Flags().IntVar(&conf.Config.Log.MaxAge, "logFileAge", 7, "log file save max days")
	configCmd.Flags().BoolVar(&conf.Config.Log.Compress, "logFileCompress", true, "compress log file")
	configCmd.Flags().BoolVar(&conf.Config.Log.EnableKafka, "logEnableKafka", false, "log send to kafka")
	viper.BindPFlag("Log.EnableKafka", configCmd.Flags().Lookup("logEnableKafka"))
	viper.BindPFlag("Log.KafkaAddress", configCmd.Flags().Lookup("kafkaAddress"))
	viper.BindPFlag("Log.FileName", configCmd.Flags().Lookup("logFile"))
	viper.BindPFlag("Log.MaxSize", configCmd.Flags().Lookup("logFileMaxSize"))
	viper.BindPFlag("Log.MaxBackups", configCmd.Flags().Lookup("logFileBackups"))
	viper.BindPFlag("Log.MaxAge", configCmd.Flags().Lookup("logFileAge"))
	viper.BindPFlag("Log.Compress", configCmd.Flags().Lookup("logFileCompress"))

	// EmailNotification
	configCmd.Flags().BoolVar(&conf.Config.EmailNotification.Enable, "emailEnable", false, "enable email send notification")
	configCmd.Flags().StringVar(&conf.Config.EmailNotification.Host, "emailHost", "smtp.qq.com", "Email host")
	configCmd.Flags().IntVar(&conf.Config.EmailNotification.Port, "emailPort", 0, "Email port")
	configCmd.Flags().StringVar(&conf.Config.EmailNotification.Username, "emailUser", "", "Email username")
	configCmd.Flags().StringVar(&conf.Config.EmailNotification.Password, "emailPw", "", "Email password")
	configCmd.Flags().StringVar(&conf.Config.EmailNotification.To, "emailTo", "", "Email to field")
	configCmd.Flags().StringVar(&conf.Config.EmailNotification.From, "emailFrom", "", "Email from field")
	configCmd.Flags().StringVar(&conf.Config.EmailNotification.Subject, "emailSubj", "", "Email subject")
	viper.BindPFlag("EmailNotification.Enable", configCmd.Flags().Lookup("emailEnable"))
	viper.BindPFlag("EmailNotification.Host", configCmd.Flags().Lookup("emailHost"))
	viper.BindPFlag("EmailNotification.Port", configCmd.Flags().Lookup("emailPort"))
	viper.BindPFlag("EmailNotification.Username", configCmd.Flags().Lookup("emailUser"))
	viper.BindPFlag("EmailNotification.Password", configCmd.Flags().Lookup("emailPw"))
	viper.BindPFlag("EmailNotification.To", configCmd.Flags().Lookup("emailTo"))
	viper.BindPFlag("EmailNotification.From", configCmd.Flags().Lookup("emailFrom"))
	viper.BindPFlag("EmailNotification.Subject", configCmd.Flags().Lookup("emailSubj"))

	// Etc
	configCmd.Flags().StringVar(&conf.Config.PidFilePath, "pid", "",
		fmt.Sprintf("GoFrameWork pid file name %s", consts.DefaultPidFilename),
	)
	configCmd.Flags().StringVar(&conf.Config.LockFilePath, "lock", "",
		fmt.Sprintf("GoFrameWork lock file name %s", consts.DefaultLockFilename),
	)
	configCmd.Flags().StringVar(&conf.Config.TempDir, "tempDir", "",
		"Temporary directory (default temporary directory of OS)")

	viper.BindPFlag("PidFilePath", configCmd.Flags().Lookup("pid"))
	viper.BindPFlag("LockFilePath", configCmd.Flags().Lookup("lock"))
	viper.BindPFlag("TempDir", configCmd.Flags().Lookup("tempDir"))

	configCmd.Flags().BoolVar(&conf.Config.TLS.Enable, "tls", false, "Enable https")
	configCmd.Flags().StringVar(&conf.Config.TLS.CertFile, "tls-cert", "", "tls private file path")
	configCmd.Flags().StringVar(&conf.Config.TLS.KeyFile, "tls-key", "", "tls public file path")

	viper.BindPFlag("TLS.Enable", configCmd.Flags().Lookup("tls"))
	viper.BindPFlag("TLS.CertFile", configCmd.Flags().Lookup("tls-cert"))
	viper.BindPFlag("TLS.KeyFile", configCmd.Flags().Lookup("tls-key"))

	configCmd.Flags().BoolVar(&conf.Config.Redis.Enable, "redis", false, "Enable Redis")
	configCmd.Flags().StringVar(&conf.Config.Redis.Addr, "redis-addr", "localhost:6379", "redis address")
	configCmd.Flags().StringVar(&conf.Config.Redis.Password, "redis-password", "", "redis password")
	configCmd.Flags().IntVar(&conf.Config.Redis.DB, "redis-db", 0, "redis database")

	viper.BindPFlag("Redis.Enable", configCmd.Flags().Lookup("redis"))
	viper.BindPFlag("Redis.Addr", configCmd.Flags().Lookup("redis-addr"))
	viper.BindPFlag("Redis.Password", configCmd.Flags().Lookup("redis-password"))
	viper.BindPFlag("Redis.DB", configCmd.Flags().Lookup("redis-db"))

	// cors
	configCmd.Flags().StringSliceVar(&conf.Config.Cors.AllowOrigins, "cors-allow-origins", []string{"*"}, "cors allow origin")
	configCmd.Flags().StringSliceVar(&conf.Config.Cors.AllowMethods, "cors-allow-methods", []string{"POST", "PUT", "GET", "DELETE", "OPTIONS"}, "cors allow origin")
	configCmd.Flags().StringSliceVar(&conf.Config.Cors.AllowHeaders, "cors-allow-headers", []string{"Authorization", "Content-Length", "X-CSRF-Token", "Accept", "Origin", "Host",
		"Connection", "Accept-Encoding", "Accept-Language", "DNT", "X-CustomHeader", "Keep-Alive", "User-Agent", "X-Requested-With",
		"If-Modified-Since", "Cache-Control", "Content-Type", "Pragma"}, "cors allow headers")
	configCmd.Flags().StringSliceVar(&conf.Config.Cors.ExposeHeaders, "cors-expose-headers", []string{"Authorization", "Content-Length", "X-CSRF-Token", "Accept", "Origin", "Host",
		"Connection", "Accept-Encoding", "Accept-Language", "DNT", "X-CustomHeader", "Keep-Alive", "User-Agent", "X-Requested-With",
		"If-Modified-Since", "Cache-Control", "Content-Type", "Pragma"}, "cors allow headers")
	configCmd.Flags().BoolVar(&conf.Config.Cors.AllowCredentials, "cors-allow-credentials", true, "allow credentials")
	configCmd.Flags().BoolVar(&conf.Config.Cors.AllowWebSockets, "cors-allow-web-sockets", true, "allow web sockets")
	configCmd.Flags().DurationVar(&conf.Config.Cors.MaxAge, "cors-max-age", 12*time.Hour, "cors max age")
	viper.BindPFlag("Cors.AllowOrigins", configCmd.Flags().Lookup("cors-allow-origins"))
	viper.BindPFlag("Cors.AllowMethods", configCmd.Flags().Lookup("cors-allow-methods"))
	viper.BindPFlag("Cors.AllowHeaders", configCmd.Flags().Lookup("cors-allow-headers"))
	viper.BindPFlag("Cors.ExposeHeaders", configCmd.Flags().Lookup("cors-expose-headers"))
	viper.BindPFlag("Cors.AllowCredentials", configCmd.Flags().Lookup("cors-allow-credentials"))
	viper.BindPFlag("Cors.AllowWebSockets", configCmd.Flags().Lookup("cors-allow-web-sockets"))
	viper.BindPFlag("Cors.MaxAge", configCmd.Flags().Lookup("cors-max-age"))

	// 初始化数据
	configCmd.Flags().BoolVar(&conf.Config.Init.Enable, "init-data", false, "Enable load data ")
	configCmd.Flags().StringVar(&conf.Config.Init.API, "init-api", "", "API data file path")

	viper.BindPFlag("Init.Enable", configCmd.Flags().Lookup("init-data"))
	viper.BindPFlag("Init.API", configCmd.Flags().Lookup("init-api"))

	// elastic
	configCmd.Flags().BoolVar(&conf.Config.Elastic.Enable, "elastic-enable", true, "Enable elastic server")
	configCmd.Flags().StringSliceVar(&conf.Config.Elastic.URLs, "elastic-urls", []string{"http://localhost:9200"}, "elastic cluster server address")
	configCmd.Flags().BoolVar(&conf.Config.Elastic.Sniff, "elastic-sniff-enable", elastic.DefaultSnifferEnabled, "Enable elastic sniff")
	configCmd.Flags().DurationVar(&conf.Config.Elastic.Sniffer, "elastic-sniffer", elastic.DefaultSnifferInterval, "sniffer interval")
	configCmd.Flags().BoolVar(&conf.Config.Elastic.HealthCheck, "elastic-health-check-enable", elastic.DefaultHealthcheckEnabled, "enable or disable healthchecks")
	configCmd.Flags().DurationVar(&conf.Config.Elastic.HealthChecker, "elastic-health-check-interval", elastic.DefaultHealthcheckInterval, "healthcheck interval")

	configCmd.Flags().StringVar(&conf.Config.Elastic.AuthUserName, "elastic-username", "", "elastic auth username")
	configCmd.Flags().StringVar(&conf.Config.Elastic.AuthPassword, "elastic-password", "", "elastic auth password")
	configCmd.Flags().BoolVar(&conf.Config.Elastic.Gzip, "elastic-gzip", elastic.DefaultGzipEnabled, "Enable elastic gzip")

	viper.BindPFlag("Elastic.Enable", configCmd.Flags().Lookup("elastic-enable"))
	viper.BindPFlag("Elastic.URLs", configCmd.Flags().Lookup("elastic-urls"))
	viper.BindPFlag("Elastic.Sniff", configCmd.Flags().Lookup("elastic-sniff-enable"))
	viper.BindPFlag("Elastic.Sniffer", configCmd.Flags().Lookup("elastic-sniffer"))
	viper.BindPFlag("Elastic.HealthCheck", configCmd.Flags().Lookup("elastic-health-check-enable"))
	viper.BindPFlag("Elastic.HealthChecker", configCmd.Flags().Lookup("elastic-health-check-interval"))

	viper.BindPFlag("Elastic.AuthUserName", configCmd.Flags().Lookup("elastic-username"))
	viper.BindPFlag("Elastic.AuthPassword", configCmd.Flags().Lookup("elastic-password"))
	viper.BindPFlag("Elastic.Gzip", configCmd.Flags().Lookup("elastic-gzip"))

}
