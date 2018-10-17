// Copyright 2018 cloudy 272685110@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package common

import "time"

//server config will used in file variables.go
type configModel struct {
	Server *serverModel `yaml:"server"`
}

//serverModel get server information from config.yml
type serverModel struct {
	Mode                 string        `yaml:"mode"`                    // run mode
	Host                 string        `yaml:"host"`                    // server host
	Port                 string        `yaml:"port"`                    // server port
	EnableHttps          bool          `yaml:"enable_https"`            // enable https
	CertFile             string        `yaml:"cert_file"`               // cert file path
	KeyFile              string        `yaml:"key_file"`                // key file path
	JwtPubKeyPath        string        `yaml:"jwt_public_key_path"`     // jwt public key path
	JwtPriKeyPath        string        `yaml:"jwt_private_key_path"`    // jwt private key path
	TokenExpireSecond    time.Duration `yaml:"token_expire_second"`     // token expire second
	SystemStaticFilePath string        `yaml:"system_static_file_path"` // system static file path
}
