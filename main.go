// Copyright 2018 cloudy 272685110@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package main

import (
	"flag"
	"fmt"
	"github.com/itcloudy/base-framework/common"
	"github.com/itcloudy/base-framework/router"
	"github.com/itcloudy/base-framework/system"
	"os"
	"path"
)

func main() {
	var err error
	fPath, _ := os.Getwd()
	fPath = path.Join(fPath, "conf")
	configPath := flag.String("c", fPath, "config file path")
	flag.Parse()
	err = system.LoadConfigInformation(*configPath)
	if err != nil {
		return
	}

	//router init
	router := router.InitRouter()
	server := common.ServerInfo
	serverInfo := common.StringsJoin(server.Host, ":", server.Port)
	// restart
	if server.EnableHttps {
		fmt.Println("server start https")
		err := router.RunTLS(serverInfo, server.CertFile, server.KeyFile)

		if err != nil {
			fmt.Println("server start failed ", err.Error())
		}
	} else {
		fmt.Printf("server start info: %s\n", serverInfo)

		err := router.Run(serverInfo)
		if err != nil {
			fmt.Println("server start failed ", err.Error())
		}
	}
}
