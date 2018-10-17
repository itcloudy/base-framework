// Copyright 2018 cloudy 272685110@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package router

import (
	"github.com/gin-gonic/gin"
	"github.com/itcloudy/base-framework/controllers"
)

func addRouter(router *gin.Engine)  {
	{
		router.GET("/", controllers.IndexGet)
	}
}
