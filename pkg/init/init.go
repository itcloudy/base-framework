// Copyright 2018 cloudy 272685110@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package init

import (
	"github.com/itcloudy/base-framework/pkg/conf"
	"github.com/itcloudy/base-framework/pkg/logs"
	"os"
)

func LoadInitData() {
	if conf.Config.Init.Enable {
		logs.Logger.Info("start load init data")
		if conf.Config.Init.API != "" {
			logs.Logger.Info("start load init api data")
			initApi()
			logs.Logger.Info("load init api data success")

		}
		if conf.Config.Init.Menu != "" {
			logs.Logger.Info("start load init menu data")
			initMenus()
			logs.Logger.Info("load init api data success")

		}
	}else{
		logs.Logger.Error("start load init data failed,config file init enable is false")
		os.Exit(-1)
	}

}
