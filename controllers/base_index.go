// Copyright 2018 cloudy 272685110@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package controllers

import "github.com/gin-gonic/gin"

func IndexGet(c *gin.Context) {

	c.Writer.Write([]byte("ok"))
}
