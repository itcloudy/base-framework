// Copyright 2018 cloudy itcloudy@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package common

import (
	"github.com/gin-gonic/gin"
	"github.com/itcloudy/base-framework/pkg/conf"
	"net/http"
)

//GenResponse generate response ,json format
func GenResponse(c *gin.Context, code int, data interface{}, message string) {
	if conf.Config.Mode == "release" {
		message = messageI18n(c.Copy(), code)
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"data":    data,
		"message": message,
	})
}
