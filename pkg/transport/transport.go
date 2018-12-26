// Copyright 2018 cloudy itcloudy@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package transport

import (
	"github.com/gin-gonic/gin"
	"github.com/itcloudy/base-framework/pkg/conf"
	"github.com/itcloudy/base-framework/pkg/logs"
	"github.com/itcloudy/base-framework/pkg/transport/restful/routers"
	"go.uber.org/zap"
)

func ServerStart() {
	if conf.Config.HTTP.Enable {
		go restStart()
	}

	if conf.Config.RPC.Enable {
		go rpcStart()
	}
	logs.Logger.Info("server start process invoke end")
}

// @title  base-framework
// @version 1.0
// @description base-framework  server
// @termsOfService https://github.com/itcloudy
func restStart() {
	router := routers.InitRestRouter()
	conf := conf.Config
	if conf.Mode == "debug" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	var err error
	if conf.TLS.Enable {
		err = router.RunTLS(conf.HTTP.Str(), conf.TLS.CertFile, conf.TLS.KeyFile)
	} else {
		err = router.Run(conf.HTTP.Str())
	}
	if err != nil {
		logs.Logger.Fatal("rest server start failed ", zap.Error(err), zap.String("http", conf.HTTP.Str()))
		panic(err)
	} else {
		logs.Logger.Info("rest server start success", zap.String("http", conf.HTTP.Str()))
	}
}

func rpcStart() {

}
