// Copyright 2018 cloudy 272685110@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package main

import (
	"github.com/itcloudy/base-framework/cmd"
	"runtime"
)

// @title  base-framework
// @version 1.0
// @description base-framework  server
// @termsOfService https://github.com/itcloudy
func main() {
	runtime.LockOSThread()
	cmd.Execute()
}
