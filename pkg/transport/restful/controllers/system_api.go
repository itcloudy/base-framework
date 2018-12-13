// Copyright 2018 cloudy 272685110@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/itcloudy/base-framework/pkg/consts"
	"github.com/itcloudy/base-framework/pkg/services"
	"github.com/itcloudy/base-framework/pkg/transport/restful/common"
	"github.com/itcloudy/base-framework/tools"
	"net/http"
)

type SystemAPIController struct {
	services.SystemAPIService
}

//根据ID获得接口详情
func (ctl SystemAPIController) CtlGetSystemAPIByID(c *gin.Context) {
	api, err := ctl.SystemAPIService.ServiceGetSystemAPIByID(c.Param("id"))
	if err != nil {

	}
	c.JSON(http.StatusOK, api)
}

//创建接口
func (ctl SystemAPIController) CtlCreateSystemAPI(c *gin.Context) {

}

//更新接口
func (ctl SystemAPIController) CtlUpdateSystemAPIByID(c *gin.Context) {

}

//查询的接口
func (ctl SystemAPIController) CtlGetAllSystemAPI(c *gin.Context) {
	var (
		page  int
		size  int
		order string
	)
	size = tools.StringToIntDefault(c.Query("size"), consts.DefaultSize)
	page = tools.StringToIntDefault(c.Query("current"), consts.DefaultPage)
	apis, _, err := ctl.SystemAPIService.ServiceGetAllSystemAPI(page, size, order, "")
	if err != nil {

	}
	common.GenResponse(c, consts.Success, apis, "")
}
