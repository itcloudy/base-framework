// +build linux freebsd darwin
// +build 386 amd64

// Copyright 2018 cloudy 272685110@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package daylight

import (
	"github.com/itcloudy/base-framework/pkg/logs"
	"github.com/itcloudy/base-framework/tools"
	"go.uber.org/zap"

	"syscall"
)

// KillPid is killing process by PID
func KillPid(pid string) error {
	err := syscall.Kill(tools.StrToInt(pid), syscall.SIGTERM)
	if err != nil {
		logs.Logger.Error("killing process failed", zap.Error(err), zap.String("pid", pid))
		return err
	}
	return nil
}
