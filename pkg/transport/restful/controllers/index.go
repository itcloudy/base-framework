// Copyright 2018 cloudy 272685110@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type IndexController struct {
}

//欢迎页面
func (ctl IndexController) CtlInformation(c *gin.Context) {
	c.JSON(http.StatusOK, "golang framework")
}

//文件上传
func (ctl IndexController) CtlFileUpload(c *gin.Context) {

}
