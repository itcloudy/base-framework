// Copyright 2018 cloudy itcloudy@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package conf

import (
	"fmt"
	"github.com/itcloudy/base-framework/pkg/consts"
	"github.com/pelletier/go-toml"
	"github.com/pkg/errors"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/viper"
)

// HostPort endpoint in form "str:int"
type HostPort struct {
	Host   string // ipaddr, hostname, or "0.0.0.0"
	Port   int    // must be in range 1..65535
	Enable bool
}

// RpcPort endpoint in form "str:int"
type RpcPort struct {
	Host   string // ipaddr, hostname, or "0.0.0.0"
	Port   int    // must be in range 1..65535
	Enable bool
}

// Str converts HostPort pair to string format
func (h HostPort) Str() string {
	return fmt.Sprintf("%s:%d", h.Host, h.Port)
}

// DBConfig database connection parameters
type DBConfig struct {
	DbType      string //mysql postgres
	Charset     string // only for mysql
	Name        string // db name
	Host        string // ipaddr, hostname, or "0.0.0.0"
	Port        int    // must be in range 1..65535
	User        string // db user
	Password    string //db password
	LockTimeout int    // lock_timeout in milliseconds
}

// CentrifugoConfig connection params
type CentrifugoConfig struct {
	Enable bool
	Secret string
	URL    string
}

func (c CentrifugoConfig) String() string {
	return fmt.Sprintf("Secret: %s URL: %s", c.Secret, c.URL)
}

// Log represents parameters of log
type LogConfig struct {
	EnableKafka  bool
	KafkaAddress []string
	FileName     string // file name
	MaxSize      int    // file size
	MaxBackups   int    // file back
	MaxAge       int    // file save days
	Compress     bool   // compress file
}

// EmailNotificationConfig smtp config
type EmailNotificationConfig struct {
	Enable   bool
	Host     string
	Port     int
	Username string
	Password string
	To       string
	From     string
	Subject  string
}
type TlsConfig struct {
	Enable   bool
	CertFile string
	KeyFile  string
}
type RedisConfig struct {
	Enable   bool
	Addr     string
	Password string
	DB       int
}
type SuperUser struct {
	UserName string
	Password string
	Email    string
	Mobile   string
}
type CorsConfig struct {
	AllowOrigins     []string
	AllowMethods     []string
	AllowHeaders     []string
	ExposeHeaders    []string
	AllowCredentials bool
	AllowWebSockets  bool
	MaxAge           time.Duration
}
type InitConfig struct {
	Enable bool   // 是否初始化
	API    string // 接口初始化路径
}
type FileUploadConfig struct {
	Target string // local,qiniu
}
type ElasticConfig struct {
	Enable        bool
	URLs          []string
	Sniff         bool
	Sniffer       time.Duration
	HealthCheck   bool
	HealthChecker time.Duration
	AuthUserName  string
	AuthPassword  string
	Gzip          bool
}
type GlobalConfig struct {
	DBUpdateToVersion string                  // 数据库升级到某个版本
	I18ns             []string                // 服务端支持的语言，用于页面和json的提示信息
	Mode              string                  // 运行模式
	TokenExpire       time.Duration           // token有效时间
	ConfigPath        string                  // 配置文件地址
	JwtPrivatePath    string                  // jwt private path
	JwtPublicPath     string                  // jwt public path
	TempDir           string                  // 临时文件
	SystemDataDir     string                  // 系统文件地址，系统自带文件
	UserDataDir       string                  // 用户文件地址，用户上传文件存储位置
	PidFilePath       string                  // 进程
	LockFilePath      string                  // lock file path
	TLS               TlsConfig               // tls
	Redis             RedisConfig             // redis
	HTTP              HostPort                // http端口
	RPC               RpcPort                 // rpc端口
	DB                DBConfig                // 数据库配置
	Centrifugo        CentrifugoConfig        // Centrifugo消息推送
	Log               LogConfig               // 日志
	EmailNotification EmailNotificationConfig // 邮件配置
	Cors              CorsConfig              // 跨域配置
	Admin             SuperUser               // 超级用户
	Init              InitConfig              // 初始化数据
	FileUpload        FileUploadConfig        // 文件上传、
	Elastic           ElasticConfig           // elastic配置

}

// GetPidPath returns path to pid file
func (c *GlobalConfig) GetPidPath() string {
	return c.PidFilePath
}

// LoadConfig from configFile
// the function has side effect updating global var Config
func LoadConfig(path string) error {
	//log.WithFields(log.Fields{"path": path}).Info("Loading config")
	return LoadConfigToVar(path, &Config)
}
func LoadConfigToVar(path string, v *GlobalConfig) error {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return errors.Errorf("Unable to load config file %s", path)
	}

	viper.SetConfigFile(path)
	err = viper.ReadInConfig()
	if err != nil {
		return errors.Wrapf(err, "reading config")
	}

	err = viper.Unmarshal(v)
	if err != nil {
		return errors.Wrapf(err, "marshalling config to global struct variable")
	}
	return nil
}

// GetConfigFromPath read config from path and returns GlobalConfig struct
func GetConfigFromPath(path string) (*GlobalConfig, error) {
	//log.WithFields(log.Fields{"path": path}).Info("Loading config")

	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return nil, errors.Errorf("Unable to load config file %s", path)
	}

	viper.SetConfigFile(path)
	err = viper.ReadInConfig()
	if err != nil {
		return nil, errors.Wrapf(err, "reading config")
	}

	c := &GlobalConfig{}
	err = viper.Unmarshal(c)
	if err != nil {
		return c, errors.Wrapf(err, "marshalling config to global struct variable")
	}

	return c, nil
}

// SaveConfig save global parameters to configFile
func SaveConfig(path string) error {
	dir := filepath.Dir(path)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.Mkdir(dir, 0775)
		if err != nil {
			return errors.Wrapf(err, "creating dir %s", dir)
		}
	}

	cf, err := os.Create(path)
	if err != nil {
		//log.WithFields(log.Fields{"type": consts.IOError, "error": err}).Error("Create config file failed")
		return err
	}
	defer cf.Close()

	err = toml.NewEncoder(cf).Encode(Config)
	if err != nil {
		return err
	}
	return nil
}

// FillRuntimePaths fills paths from runtime parameters
func FillRuntimePaths() error {
	cwd, err := os.Getwd()
	if err != nil {
		return errors.Wrapf(err, "getting current wd")
	}
	if Config.SystemDataDir == "" {
		Config.SystemDataDir = filepath.Join(cwd, consts.DefaultSystemDataDirName)
	}
	if Config.UserDataDir == "" {
		Config.UserDataDir = filepath.Join(cwd, consts.DefaultUserDataDirName)
	}

	if Config.TempDir == "" {
		Config.TempDir = filepath.Join(cwd, consts.DefaultTempDirName)
	}

	if Config.PidFilePath == "" {
		Config.PidFilePath = filepath.Join(Config.SystemDataDir, consts.DefaultPidFilename)
	}
	if Config.LockFilePath == "" {
		Config.LockFilePath = filepath.Join(Config.SystemDataDir, consts.DefaultLockFilename)
	}
	if Config.Log.FileName == "" {
		Config.Log.FileName = filepath.Join(cwd, "logs", consts.DefaultLogFileName)

	}
	//jwt
	if Config.JwtPrivatePath == "" {
		Config.JwtPrivatePath = filepath.Join(cwd, "config", "jwt", "tm.rsa")
	}
	if Config.JwtPublicPath == "" {
		Config.JwtPublicPath = filepath.Join(cwd, "config", "jwt", "tm.rsa.pub")
	}
	//https
	if Config.TLS.CertFile == "" {
		Config.TLS.CertFile = filepath.Join(cwd, "config", "https", "cert.pem")
	}
	if Config.TLS.KeyFile == "" {
		Config.TLS.KeyFile = filepath.Join(cwd, "config", "https", "key.pem")
	}

	if Config.Init.API == "" {
		Config.Init.API = filepath.Join(cwd, "init", "api_data.yml")

	}
	return nil
}
