// Copyright 2018 cloudy itcloudy@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package routers

import (
	"github.com/casbin/casbin"
	"github.com/casbin/gorm-adapter"
	"github.com/gin-gonic/gin"
	"github.com/itcloudy/base-framework/pkg/conf"
	"github.com/itcloudy/base-framework/pkg/consts"
	"github.com/itcloudy/base-framework/pkg/logs"
	"github.com/itcloudy/base-framework/pkg/transport/restful/middles"
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
	connectInfo := conf.GetDBConnectionString(dbConf.DbType, dbConf.Host, dbConf.Port, dbConf.User, dbConf.Password, dbConf.Name, dbConf.Charset)
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
		//欢迎页面
		router.GET("/", rest.IndexContainer().CtlInformation)
		//文件上传
		router.POST("/file/upload", rest.IndexContainer().CtlFileUpload)
		//登录
		router.POST("/login/account", rest.UserContainer().CtlLoginAccount)

	}

	// auth router
	authRouter := router.Group("/auth")
	authRouter.Use(middles.CasbinJwtAuthorize(conf.Enforcer))
	{
		//获得个人信息
		authRouter.GET("/self_info", rest.UserContainer().CtlGetSelfInformation)

		//管理员根据ID获得某个用户信息
		authRouter.GET("/user/:id", rest.UserContainer().CtlGetUserByID)
		//管理员根据用户名获得某个用户信息
		authRouter.GET("/username/:username", rest.UserContainer().CtlGetUserByUserName)

		//角色操作
		authRouter.GET("/role/:id", rest.RoleContainer().CtlGetRoleByID)
		authRouter.POST("/role", rest.RoleContainer().CtlCreateRole)
		authRouter.PUT("/role/:id", rest.RoleContainer().CtlUpdateRoleByID)
		authRouter.GET("/roles", rest.RoleContainer().CtlGetAllRole)
		//系统接口操作
		authRouter.GET("/system_api/:id", rest.SystemAPIContainer().CtlGetSystemAPIByID)
		authRouter.POST("/system_api_active", rest.SystemAPIContainer().CtlActiveActionSystemAPI)
		authRouter.POST("/system_api", rest.SystemAPIContainer().CtlCreateSystemAPI)
		authRouter.PUT("/system_api/:id", rest.SystemAPIContainer().CtlUpdateSystemAPIByID)
		authRouter.GET("/system_apis", rest.SystemAPIContainer().CtlGetAllSystemAPI)
	}

}
