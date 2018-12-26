// Copyright 2018 cloudy itcloudy@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package daylight

import (
	"github.com/itcloudy/base-framework/pkg/logs"
	"github.com/itcloudy/base-framework/pkg/transport"
	"github.com/itcloudy/base-framework/pkg/transport/restful/middles"
	"go.uber.org/zap"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"

	"github.com/itcloudy/base-framework/tools"

	"github.com/itcloudy/base-framework/pkg/conf"
	"github.com/itcloudy/base-framework/pkg/consts"
)

// Start starts the main code of the program
func Start() {

	defer func() {

		if r := recover(); r != nil {
			logs.Logger.Fatal("recovered panic", zap.String("type", consts.PanicRecoveredError))
			panic(r)
		}
	}()
	logs.Logger.Sync()

	Exit := func(code int) {
		delPidFile()
		os.Exit(code)
	}

	f := tools.LockOrDie(conf.Config.LockFilePath)
	defer f.Unlock()

	if err := tools.MakeDirectory(conf.Config.TempDir); err != nil {
		logs.Logger.Fatal("can't create temporary directory", zap.String("type", consts.IOError), zap.String("dir", conf.Config.TempDir))
		Exit(1)
	}

	killOld()

	rand.Seed(time.Now().UTC().UnixNano())

	// save the current pid and version
	if err := savePid(); err != nil {
		logs.Logger.Fatal("can't create pid", zap.Error(err))

		Exit(1)
	}
	defer delPidFile()
	//加载jwt使用的公私钥
	middles.InitKeys()
	// 加载语言资源
	conf.BundleI18nLanguages()
	cfg := conf.Config.DB

	conf.GetDBConnection(cfg.DbType, cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name, cfg.Charset, "")
	if conf.Config.Elastic.Enable {
		conf.GetElasticClient(conf.Config.Elastic)
	}
	transport.ServerStart()

	select {}
}

func delPidFile() {
	os.Remove(conf.Config.GetPidPath())
}

func killOld() {
	pidPath := conf.Config.GetPidPath()
	if _, err := os.Stat(pidPath); err == nil {
		dat, err := ioutil.ReadFile(pidPath)
		if err != nil {
			logs.Logger.Error("reading pid file failed", zap.String("path", pidPath), zap.Error(err))
		}
		var pidMap map[string]string
		err = json.Unmarshal(dat, &pidMap)
		if err != nil {
			logs.Logger.Error("un marshalling pid map", zap.String("type", consts.JSONUnmarshallError), zap.ByteString("data", dat), zap.Error(err))

		}
		logs.Logger.Debug("old pid path", zap.String("path", pidPath))

		KillPid(pidMap["pid"])
		if fmt.Sprintf("%s", err) != "null" {
			// give 15 sec to end the previous process
			for i := 0; i < 15; i++ {
				if _, err := os.Stat(conf.Config.GetPidPath()); err == nil {
					time.Sleep(time.Second)
				} else {
					break
				}
			}
		}
	}
}

func savePid() error {
	pid := os.Getpid()
	PidAndVer, err := json.Marshal(map[string]string{"pid": tools.IntToStr(pid), "version": consts.VERSION})
	if err != nil {
		logs.Logger.Error("marshalling pid to json", zap.String("type", consts.JSONMarshallError), zap.Error(err), zap.Int("pid", pid))
		return err
	}

	return ioutil.WriteFile(conf.Config.GetPidPath(), PidAndVer, 0644)
}
