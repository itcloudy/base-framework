// +build windows

// Copyright 2018 cloudy 272685110@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package daylight

import (
	"fmt"

	"os/exec"
	"regexp"
)

// KillPid kills the process with the specified pid
func KillPid(pid string) error {

	rez, err := exec.Command("tasklist", "/fi", "PID eq "+pid).Output()
	if err != nil {
		log.WithFields(log.Fields{"type": consts.CommandExecutionError, "err": err, "cmd": "tasklist /fi PID eq" + pid}).Error("Error executing command")
		return err
	}
	if string(rez) == "" {
		return fmt.Errorf("null")
	}
	log.WithFields(log.Fields{"cmd": "tasklist /fi PID eq " + pid}).Debug("command execution result")
	if ok, _ := regexp.MatchString(`(?i)PID`, string(rez)); !ok {
		return fmt.Errorf("null")
	}
	return nil
}
