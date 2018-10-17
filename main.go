// Copyright 2018 cloudy 272685110@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package main

import (
	"flag"
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
}
