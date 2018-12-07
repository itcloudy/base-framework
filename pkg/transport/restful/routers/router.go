// Copyright 2018 cloudy 272685110@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/itcloudy/base-framework/pkg/conf"
	"github.com/itcloudy/base-framework/pkg/consts"
	"github.com/itcloudy/base-framework/pkg/transport/restful/middles"
	"time"
)

func setTemplate(router *gin.Engine) {
	//user upload file dir
	router.StaticFS(consts.USER_UPLOAD_FILE_URL, gin.Dir(consts.DefaultUserDataDirName, false))
	// system file dir
	router.StaticFS(consts.SYSTEM_STATIC_FILE_URL, gin.Dir(consts.DefaultSystemDataDirName, false))
}

//InitRestRouter all rest api
func InitRestRouter() *gin.Engine {
	router := gin.Default()
	//cors
	Cors := conf.Config.Cors
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowCredentials = Cors.AllowCredentials
	corsConfig.AllowHeaders = Cors.AllowHeaders
	corsConfig.AllowMethods = Cors.AllowMethods
	corsConfig.AllowOrigins = Cors.AllowOrigins
	corsConfig.AllowWebSockets = Cors.AllowWebSockets
	corsConfig.ExposeHeaders = Cors.ExposeHeaders
	corsConfig.MaxAge = Cors.MaxAge * time.Hour
	router.Use(cors.New(corsConfig))
	// set template
	setTemplate(router)
	// jwt auth
	router.Use(middles.JwtAuthorize())
	// recovery
	router.Use(gin.Recovery())
	// add rest api
	addRouter(router)

	return router

}
