// Copyright 2018 cloudy 272685110@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package router

import (
	"github.com/gin-gonic/gin"
)

//InitRouter router init
func InitRouter() *gin.Engine {

	router := gin.Default()


	router.Use(gin.Recovery())
	// add routers
	addRouter(router)
	return router
}
