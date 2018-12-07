// Copyright 2018 cloudy 272685110@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package routers

import (
	"fmt"
	"github.com/casbin/casbin"
	"github.com/casbin/gorm-adapter"
	"github.com/gin-gonic/gin"
	"github.com/itcloudy/base-framework/pkg/conf"
	"github.com/itcloudy/base-framework/pkg/consts"
	"github.com/itcloudy/base-framework/pkg/logs"
	"github.com/itcloudy/base-framework/pkg/models"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"go.uber.org/zap"
	"os"
	"path"
)

func addRouter(router *gin.Engine) {
	defer func() {

		if r := recover(); r != nil {
			logs.Logger.Fatal("rest router init failed")
			panic(r)
		}
	}()
	// swagger docs
	if conf.Config.Mode == "debug" {
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	// casbin load role permission
	var adapter *gormadapter.Adapter
	dbConf := conf.Config.DB
	connectInfo := models.GetDBConnectionString(dbConf.DbType, dbConf.Host, dbConf.Port, dbConf.User, dbConf.Password, dbConf.Name, dbConf.Charset)
	fmt.Println(connectInfo)
	adapter = gormadapter.NewAdapter(dbConf.DbType, connectInfo, true)
	cwd, _ := os.Getwd()
	conf.Enforcer = casbin.NewEnforcer(path.Join(cwd, consts.DefaultWorkdirName, "casbin_rbac_model.conf"), adapter)
	err := conf.Enforcer.LoadPolicy()
	if err != nil {
		logs.Logger.Fatal("can't init router", zap.Error(err))

	}

	// 初始化所有表对象对应的controller，根据配置注入对应的数据库连接信息等
	rest := restContainer()
	// public router
	{
	}

	// auth router
	authRouter := router.Group("/auth")
	//authRouter.Use(middles.CasbinJwtAuthorize(conf.Enforcer))
	{
		authRouter.GET("/user/:id", rest.UserContainer().CtlGetUserByID)
		authRouter.GET("/username/:username", rest.UserContainer().CtlGetUserByUserName)
		//authRouter.GET("/menu/self", rest.MenuContainer().CtlGetSelfMenu)

	}

}
