// Copyright 2018 cloudy 272685110@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package system

import (
	"fmt"
	"github.com/itcloudy/base-framework/common"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path"
)

//LoadConfigInformation load config information for application
func LoadConfigInformation(configPath string) (err error) {
	var (
		filePath string
		wr       string
	)

	if configPath == "" {
		wr, _ = os.Getwd()
		wr = path.Join(wr, "conf")

	} else {
		wr = configPath
	}
	common.WorkSpace = wr
	filePath = path.Join(common.WorkSpace, "config.yml")
	configData, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf(" config file read failed: %s", err)
		os.Exit(-1)

	}
	err = yaml.Unmarshal(configData, &common.ConfigInfo)
	if err != nil {
		fmt.Printf(" config parse failed: %s", err)

		os.Exit(-1)
	}
	// server information
	common.ServerInfo = common.ConfigInfo.Server
	return nil
}
